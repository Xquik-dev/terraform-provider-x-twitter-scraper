package x_write

import (
	"context"
	"fmt"

	"github.com/Xquik-dev/x-twitter-scraper-go"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithConfigure = (*writeResource)(nil)

type writeResource struct {
	client    *xtwitterscraper.Client
	operation operation
}

func newResource(op operation) resource.Resource {
	return &writeResource{operation: op}
}

func resourceFor(action string) resource.Resource {
	for _, op := range operations {
		if op.action == action {
			return newResource(op)
		}
	}
	panic("unknown canonical write action: " + action)
}

func NewCreateTweetResource() resource.Resource     { return resourceFor("create_tweet") }
func NewDeleteTweetResource() resource.Resource     { return resourceFor("delete_tweet") }
func NewLikeResource() resource.Resource            { return resourceFor("like") }
func NewUnlikeResource() resource.Resource          { return resourceFor("unlike") }
func NewRetweetResource() resource.Resource         { return resourceFor("retweet") }
func NewUnretweetResource() resource.Resource       { return resourceFor("unretweet") }
func NewFollowResource() resource.Resource          { return resourceFor("follow") }
func NewUnfollowResource() resource.Resource        { return resourceFor("unfollow") }
func NewRemoveFollowerResource() resource.Resource  { return resourceFor("remove_follower") }
func NewSendDMResource() resource.Resource          { return resourceFor("send_dm") }
func NewUploadMediaResource() resource.Resource     { return resourceFor("upload_media") }
func NewUpdateProfileResource() resource.Resource   { return resourceFor("update_profile") }
func NewUpdateAvatarResource() resource.Resource    { return resourceFor("update_avatar") }
func NewUpdateBannerResource() resource.Resource    { return resourceFor("update_banner") }
func NewCreateCommunityResource() resource.Resource { return resourceFor("create_community") }
func NewDeleteCommunityResource() resource.Resource { return resourceFor("delete_community") }
func NewJoinCommunityResource() resource.Resource   { return resourceFor("join_community") }
func NewLeaveCommunityResource() resource.Resource  { return resourceFor("leave_community") }

func (r *writeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.operation.resourceSuffix
}

func (r *writeResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resourceSchema(ctx, r.operation)
}

func (r *writeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*xtwitterscraper.Client)
	if !ok {
		resp.Diagnostics.AddError("unexpected resource configure type", fmt.Sprintf("Expected *xtwitterscraper.Client, got %T.", req.ProviderData))
		return
	}
	r.client = client
}

func (r *writeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data resourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	action, err := executeWrite(ctx, r.client, r.operation, writeRequest{
		account:        data.Account.ValueString(),
		idempotencyKey: data.IdempotencyKey.ValueString(),
		targetID:       data.TargetID.ValueString(),
		payloadJSON:    data.PayloadJSON.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("write did not reach a verified terminal state", err.Error())
		return
	}
	data.apply(action)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if action.Terminal && !action.Success {
		resp.Diagnostics.AddError("write reached a terminal failure", fmt.Sprintf("Action %s ended with status %s. safeToRetry=%t. No automatic retry was attempted.", action.WriteActionID, action.Status, action.SafeToRetry))
	}
}

func (r *writeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data resourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	action := writeAction{WriteActionID: data.WriteActionID.ValueString(), StatusURL: data.StatusURL.ValueString()}
	path, err := canonicalPollPath(action.StatusURL, action.WriteActionID)
	if err != nil {
		resp.Diagnostics.AddError("invalid write state", err.Error())
		return
	}
	if err := r.client.Get(ctx, path, nil, &action); err != nil {
		resp.Diagnostics.AddError("failed to refresh write action", err.Error())
		return
	}
	if err := validateAction(r.operation, action); err != nil {
		resp.Diagnostics.AddError("invalid canonical write response", err.Error())
		return
	}
	if !action.Terminal {
		action, err = pollUntilTerminal(ctx, r.client, r.operation, action)
		if err != nil {
			resp.Diagnostics.AddError("write did not reach a verified terminal state", err.Error())
			return
		}
	}
	data.apply(action)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *writeResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("immutable write action", "Changing a write action must replace the resource and requires a new Idempotency-Key.")
}

func (r *writeResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// A write action is an immutable audit record. Destroy removes only Terraform state.
}
