// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_tweet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*XTweetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"account": schema.StringAttribute{
				Description:   "X account (@username or account ID)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"attachment_url": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"community_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"is_note_tweet": schema.BoolAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"reply_to_tweet_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"text": schema.StringAttribute{
				Description:   "Tweet text (optional when media is provided)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"media": schema.ListAttribute{
				Description:   "Array of media URLs to attach (mutually exclusive with media_ids)",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"media_ids": schema.ListAttribute{
				Description:   "Array of media IDs to attach (mutually exclusive with media)",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"success": schema.BoolAttribute{
				Computed: true,
			},
			"tweet_id": schema.StringAttribute{
				Computed: true,
			},
			"author": schema.SingleNestedAttribute{
				Description: "Author of a tweet with follower count and verification status.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetAuthorModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"followers": schema.Int64Attribute{
						Computed: true,
					},
					"username": schema.StringAttribute{
						Computed: true,
					},
					"verified": schema.BoolAttribute{
						Computed: true,
					},
					"profile_picture": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"tweet": schema.SingleNestedAttribute{
				Description: "Full tweet with text, engagement metrics, media, and metadata.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[XTweetTweetModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"bookmark_count": schema.Int64Attribute{
						Computed: true,
					},
					"like_count": schema.Int64Attribute{
						Computed: true,
					},
					"quote_count": schema.Int64Attribute{
						Computed: true,
					},
					"reply_count": schema.Int64Attribute{
						Computed: true,
					},
					"retweet_count": schema.Int64Attribute{
						Computed: true,
					},
					"text": schema.StringAttribute{
						Computed: true,
					},
					"view_count": schema.Int64Attribute{
						Computed: true,
					},
					"conversation_id": schema.StringAttribute{
						Description: "ID of the root tweet in the conversation thread",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"entities": schema.MapAttribute{
						Description: "Parsed entities from the tweet text (URLs, mentions, hashtags, media)",
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
					},
					"is_note_tweet": schema.BoolAttribute{
						Description: "Whether this is a Note Tweet (long-form post, up to 25,000 characters)",
						Computed:    true,
					},
					"is_quote_status": schema.BoolAttribute{
						Description: "Whether this tweet quotes another tweet",
						Computed:    true,
					},
					"is_reply": schema.BoolAttribute{
						Description: "Whether this tweet is a reply to another tweet",
						Computed:    true,
					},
					"media": schema.ListNestedAttribute{
						Description: "Attached media items, omitted when the tweet has no media",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[XTweetTweetMediaModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"media_url": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Description: `Available values: "photo", "video", "animated_gif".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"photo",
											"video",
											"animated_gif",
										),
									},
								},
								"url": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"quoted_tweet": schema.MapAttribute{
						Description: "The quoted tweet object, present when isQuoteStatus is true",
						Computed:    true,
						CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
						ElementType: jsontypes.NormalizedType{},
					},
					"source": schema.StringAttribute{
						Description: "Client application used to post this tweet",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *XTweetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *XTweetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
