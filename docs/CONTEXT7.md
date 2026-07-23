# Context7 Guide

Use this guide when an agent needs the provider contract quickly.

The provider manages monitors, webhooks, content workflows, and X write actions.

Use the REST API for tweet search, timelines, followers, and export jobs.

## Install

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source  = "Xquik-dev/x-twitter-scraper"
      version = "~> 0.3.1"
    }
  }
}
```

Run `terraform init`.

## Authenticate

Prefer environment variables:

```sh
export X_TWITTER_SCRAPER_API_KEY="your-api-key"
```

Bearer authentication uses `X_TWITTER_SCRAPER_BEARER_TOKEN`.

Never commit credentials, Terraform state, logs, or local plugin configuration.

## Configure

```hcl
provider "x-twitter-scraper" {}
```

Run `terraform validate` before applying changes.

## Monitor and Deliver Events

```hcl
resource "x-twitter-scraper_monitor" "product_updates" {
  username    = "xquik"
  event_types = ["tweet.new"]
}

resource "x-twitter-scraper_webhook" "events" {
  url         = "https://example.com/xquik/webhook"
  event_types = ["tweet.new"]
}
```

Treat the returned webhook HMAC secret as sensitive.

## Create a Durable Write

```hcl
resource "x-twitter-scraper_x_tweet" "announcement" {
  account         = "@example"
  idempotency_key = "terraform-announcement-v1"
  payload_json = jsonencode({
    text = "Published through the Xquik Terraform provider."
  })
}
```

Use one stable idempotency key for each exact intended request.

Changing a write replaces the resource and requires a new key.

The provider waits for a verified terminal write state.

It never blindly retries an uncertain dispatch.

## Resource Map

- Content: `compose`, `draft`, `style`
- Monitoring: `monitor`, `monitor_keyword`, `webhook`
- Posts: `x_tweet`, `x_tweet_delete`
- Engagement: `x_tweet_like`, `x_tweet_unlike`, `x_tweet_retweet`, `x_tweet_unretweet`
- Relationships: `x_user_follow`, `x_user_unfollow`, `x_user_remove_follower`
- Messaging and media: `x_dm`, `x_media`
- Profiles: `x_profile`, `x_profile_avatar`, `x_profile_banner`
- Communities: `x_community`, `x_community_delete`, `x_community_join`, `x_community_leave`
- Support and checkout: `support_ticket`, `guest_wallet`

Prefix each name with `x-twitter-scraper_` inside Terraform blocks.

## Data Source Map

- Usage: `account`
- Content: `draft`, `style`
- Monitoring: `monitor`, `monitor_keyword`, `event`
- Draws: `draw`
- X lookups: `x_tweet`, `x_user`, `x_account`
- Write status: `x_write_action`
- Support: `support_ticket`

Data sources read existing state.

They do not provide tweet search or timeline scraping.

## Public References

- [Quickstart](guides/quickstart.md)
- [Generated resources](resources)
- [Generated data sources](data-sources)
- [Examples](../examples)
- [REST API](https://docs.xquik.com/api-reference/overview)
- [OpenAPI](https://xquik.com/openapi.json)
- [Security policy](../SECURITY.md)

## Agent Rules

- Read generated schema pages before writing Terraform.
- Never invent fields from REST request bodies.
- Preserve `payload_json` as an operation-specific JSON object.
- Never place `account` inside `payload_json`.
- Never replace idempotency keys during uncertain retries.
- Treat plans, state, credentials, and webhook secrets as sensitive.
- Use only documented resource and data source names.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
