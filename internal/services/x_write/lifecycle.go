package x_write

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/Xquik-dev/x-twitter-scraper-go/option"
)

const maxPollDuration = 15 * time.Minute

var terminalStatuses = map[string]bool{
	"failed":  true,
	"expired": true,
	"success": true,
}

var validStatuses = map[string]bool{
	"accepted":             true,
	"dispatching":          true,
	"pending_confirmation": true,
	"success":              true,
	"failed":               true,
	"expired":              true,
}

var validBillingStatuses = map[string]bool{
	"not_charged":   true,
	"pending":       true,
	"charged":       true,
	"charge_failed": true,
	"refunded":      true,
}

type writeRequest struct {
	account        string
	idempotencyKey string
	targetID       string
	payloadJSON    string
}

type writeAction struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	WriteActionID  string `json:"writeActionId"`
	Action         string `json:"action"`
	Status         string `json:"status"`
	Terminal       bool   `json:"terminal"`
	Retryable      bool   `json:"retryable"`
	SafeToRetry    bool   `json:"safeToRetry"`
	StatusURL      string `json:"statusUrl"`
	PollAfterMs    int64  `json:"pollAfterMs"`
	Charged        bool   `json:"charged"`
	ChargedCredits string `json:"chargedCredits"`
	TargetID       string `json:"targetId"`
	RequestID      string `json:"requestId"`
	SendDispatched bool   `json:"sendDispatched"`
	Success        bool   `json:"success"`
	Error          string `json:"error"`
	Message        string `json:"message"`
	Billing        struct {
		Charged        bool   `json:"charged"`
		ChargedCredits string `json:"chargedCredits"`
		PlannedCredits string `json:"plannedCredits"`
		Status         string `json:"status"`
	} `json:"billing"`
	Request struct {
		Hash    string          `json:"hash"`
		Payload json.RawMessage `json:"payload"`
	} `json:"request"`
	Account struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"account"`
	Result struct {
		ID    string `json:"id"`
		State string `json:"state"`
		Type  string `json:"type"`
	} `json:"result"`
	Target struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"target"`
	NextAction struct {
		Type                      string `json:"type"`
		AfterMs                   int64  `json:"afterMs"`
		RequiresNewIdempotencyKey bool   `json:"requiresNewIdempotencyKey"`
		URL                       string `json:"url"`
	} `json:"nextAction"`
}

func executeWrite(ctx context.Context, client *xtwitterscraper.Client, op operation, request writeRequest) (writeAction, error) {
	if strings.TrimSpace(request.account) == "" {
		return writeAction{}, fmt.Errorf("account must not be empty")
	}
	if !validIdempotencyKey(request.idempotencyKey) {
		return writeAction{}, fmt.Errorf("Idempotency-Key must contain 1-255 visible ASCII characters")
	}
	body, err := op.requestBody(request.account, request.payloadJSON)
	if err != nil {
		return writeAction{}, err
	}
	if op.targetRequired && strings.TrimSpace(request.targetID) == "" {
		return writeAction{}, fmt.Errorf("target_id is required for %s", op.action)
	}

	var action writeAction
	err = client.Execute(
		ctx,
		op.method,
		op.path(request.targetID),
		body,
		&action,
		option.WithHeader("Idempotency-Key", request.idempotencyKey),
		option.WithMaxRetries(0),
	)
	if err != nil {
		return writeAction{}, fmt.Errorf("write dispatch result is unknown; retry only the exact request with the same Idempotency-Key or verify its result: %w", err)
	}
	if err := validateAction(op, action); err != nil {
		return writeAction{}, err
	}
	if action.Terminal {
		return action, nil
	}

	pollCtx := ctx
	cancel := func() {}
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		pollCtx, cancel = context.WithTimeout(ctx, maxPollDuration)
	}
	defer cancel()

	return pollUntilTerminal(pollCtx, client, op, action)
}

func pollUntilTerminal(ctx context.Context, client *xtwitterscraper.Client, op operation, action writeAction) (writeAction, error) {
	for !action.Terminal {
		path, err := canonicalPollPath(action.StatusURL, action.WriteActionID)
		if err != nil {
			return action, err
		}
		if err := waitForPoll(ctx, action.PollAfterMs); err != nil {
			return action, fmt.Errorf("write %s is still non-terminal; poll %s before retrying: %w", action.WriteActionID, action.StatusURL, err)
		}
		if err := client.Get(
			ctx,
			path,
			nil,
			&action,
		); err != nil {
			return action, fmt.Errorf("failed to poll write %s; do not redispatch while dispatch state is unknown: %w", action.WriteActionID, err)
		}
		if err := validateAction(op, action); err != nil {
			return action, err
		}
	}
	return action, nil
}

func validateAction(op operation, action writeAction) error {
	if action.Object != "x_write_action" {
		return fmt.Errorf("canonical write response has unexpected object %q", action.Object)
	}
	if action.WriteActionID == "" || action.ID == "" {
		return fmt.Errorf("canonical write response omitted id or writeActionId")
	}
	if action.ID != action.WriteActionID {
		return fmt.Errorf("canonical write response id and writeActionId differ")
	}
	if action.Action != op.action {
		return fmt.Errorf("canonical write response action %q does not match %q", action.Action, op.action)
	}
	if !validStatuses[action.Status] {
		return fmt.Errorf("canonical write response has unknown status %q", action.Status)
	}
	if action.Terminal != terminalStatuses[action.Status] {
		return fmt.Errorf("canonical write response status %q and terminal=%t disagree", action.Status, action.Terminal)
	}
	if !validBillingStatuses[action.Billing.Status] {
		return fmt.Errorf("canonical write response has unknown billing status %q", action.Billing.Status)
	}
	if action.Charged != action.Billing.Charged || action.ChargedCredits != action.Billing.ChargedCredits {
		return fmt.Errorf("canonical write response billing compatibility fields disagree")
	}
	if action.Target.ID != "" && action.TargetID != action.Target.ID {
		return fmt.Errorf("canonical write response target compatibility fields disagree")
	}
	if action.SafeToRetry && action.SendDispatched {
		return fmt.Errorf("canonical write response marks a dispatched write safe to retry")
	}
	if action.NextAction.RequiresNewIdempotencyKey && !action.SafeToRetry {
		return fmt.Errorf("canonical write response requires a new Idempotency-Key when safeToRetry is false")
	}
	if action.Success != (action.Status == "success") {
		return fmt.Errorf("canonical write response status %q and success=%t disagree", action.Status, action.Success)
	}
	if action.StatusURL == "" {
		return fmt.Errorf("canonical write response for %s omitted statusUrl", action.WriteActionID)
	}
	return nil
}

func validIdempotencyKey(key string) bool {
	if len(key) < 1 || len(key) > 255 {
		return false
	}
	for index := 0; index < len(key); index++ {
		if key[index] < '!' || key[index] > '~' {
			return false
		}
	}
	return true
}

func canonicalPollPath(statusURL string, writeActionID string) (string, error) {
	parsed, err := url.Parse(statusURL)
	if err != nil {
		return "", fmt.Errorf("invalid statusUrl for write %s: %w", writeActionID, err)
	}
	expectedSuffix := "/x/write-actions/" + url.PathEscape(writeActionID)
	if !strings.HasSuffix(parsed.EscapedPath(), expectedSuffix) {
		return "", fmt.Errorf("statusUrl for write %s is not its canonical polling URL", writeActionID)
	}
	return strings.TrimPrefix(expectedSuffix, "/"), nil
}

func waitForPoll(ctx context.Context, pollAfterMs int64) error {
	if pollAfterMs < 0 {
		pollAfterMs = 0
	}
	maxPollMilliseconds := int64(maxPollDuration / time.Millisecond)
	if pollAfterMs > maxPollMilliseconds {
		pollAfterMs = maxPollMilliseconds
	}
	timer := time.NewTimer(time.Duration(pollAfterMs) * time.Millisecond)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
