// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package style

import (
	"github.com/Xquik-dev/terraform-provider-x-twitter-scraper/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StyleModel struct {
	ID           types.String         `tfsdk:"id" path:"id,required"`
	Label        types.String         `tfsdk:"label" json:"label" api:"required,no_refresh"`
	Tweets       *[]*StyleTweetsModel `tfsdk:"tweets" json:"tweets" api:"required"`
	FetchedAt    timetypes.RFC3339    `tfsdk:"fetched_at" json:"fetchedAt" api:"computed" format:"date-time"`
	IsOwnAccount types.Bool           `tfsdk:"is_own_account" json:"isOwnAccount" api:"computed"`
	TweetCount   types.Int64          `tfsdk:"tweet_count" json:"tweetCount" api:"computed"`
	XUsername    types.String         `tfsdk:"x_username" json:"xUsername" api:"computed"`
}

func (m StyleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StyleModel) MarshalJSONForUpdate(state StyleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StyleTweetsModel struct {
	Text types.String `tfsdk:"text" json:"text" api:"required"`
}
