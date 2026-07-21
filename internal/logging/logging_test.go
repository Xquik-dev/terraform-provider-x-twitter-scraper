package logging

import "testing"

func TestRedactHeader(t *testing.T) {
	tests := map[string]string{
		"Authorization":   "[REDACTED]",
		"Cookie":          "[REDACTED]",
		"Idempotency-Key": "[REDACTED]",
		"Set-Cookie":      "[REDACTED]",
		"X-API-Key":       "[REDACTED]",
		"Content-Type":    "application/json",
	}
	for name, want := range tests {
		value := "secret"
		if name == "Content-Type" {
			value = want
		}
		if got := redactHeader(name, value); got != want {
			t.Errorf("redactHeader(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestRedactJSONBodyRecursively(t *testing.T) {
	body := []byte(`{"api_key":"key","nested":{"accessToken":"token","safe":"value"},"items":[{"password":"password"}]}`)
	want := `{"api_key":"[REDACTED]","items":[{"password":"[REDACTED]"}],"nested":{"accessToken":"[REDACTED]","safe":"value"}}`
	if got := redactJSONBody(body); got != want {
		t.Errorf("redactJSONBody() = %q, want %q", got, want)
	}
}

func TestRedactJSONBodyOmitsNonJSON(t *testing.T) {
	const body = "plain text"
	if got := redactJSONBody([]byte(body)); got != omittedBody {
		t.Errorf("redactJSONBody() = %q, want %q", got, omittedBody)
	}
}

func TestRedactJSONBodyOmitsEmptyBody(t *testing.T) {
	if got := redactJSONBody(nil); got != omittedBody {
		t.Errorf("redactJSONBody() = %q, want %q", got, omittedBody)
	}
}
