# X API & Twitter API Terraform provider by Xquik

[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13729/badge)](https://www.bestpractices.dev/projects/13729)

Use this X API & Twitter API Terraform provider to manage X automation as code.
It turns Xquik monitors, signed webhooks, drafts & X actions into Terraform resources.
Scrape Twitter search, timelines & followers with the [Xquik REST API](https://docs.xquik.com/api-reference/overview).
That API works as a Twitter scraper & tweet scraper for your apps.
The provider is a Twitter API alternative built on that same API.

Xquik is the world's fastest and cheapest X Twitter scraper service with the most complete X data.
Xquik charges only for delivered, unique, filter-matching results and never pads rows.

> Xquik is an independent third-party service. Not affiliated with X Corp.
> "Twitter" and "X" are trademarks of X Corp.

[Quickstart](docs/guides/quickstart.md) | [Provider docs](docs/index.md) | [Examples](examples) | [REST API](https://docs.xquik.com/api-reference/overview) | [OpenAPI](https://xquik.com/openapi.json) | [Security](SECURITY.md)

## Manage durable X automation

| Workflow | Resources |
| --- | --- |
| Compose and refine posts | `compose`, `draft`, `style` |
| Monitor accounts and keywords | `monitor`, `monitor_keyword` |
| Deliver signed events | `webhook` |
| Publish and delete posts | `x_tweet`, `x_tweet_delete` |
| Like, unlike, repost, and undo reposts | `x_tweet_like`, `x_tweet_unlike`, `x_tweet_retweet`, `x_tweet_unretweet` |
| Follow, unfollow, and remove followers | `x_user_follow`, `x_user_unfollow`, `x_user_remove_follower` |
| Send direct messages and upload media | `x_dm`, `x_media` |
| Update profiles, avatars, and banners | `x_profile`, `x_profile_avatar`, `x_profile_banner` |
| Create, delete, join, and leave communities | `x_community`, `x_community_delete`, `x_community_join`, `x_community_leave` |
| Request support and guest checkout | `support_ticket`, `guest_wallet` |

Every X write resource stores its canonical response. Uncertain retries keep the same idempotency key.

## Read existing Xquik data

- `account` for account usage and limits
- `draft` and `style` for drafts and writing styles
- `monitor` and `monitor_keyword` for account and keyword monitors
- `event` and `draw` for events and giveaway draws
- `x_tweet`, `x_user` and `x_account` for tweet, user and account lookups
- `x_write_action` for durable write status
- `support_ticket` for support requests

## Run Twitter search outside Terraform

Use the [Xquik REST API](https://docs.xquik.com/api-reference/overview) for Twitter search, timelines, follower exports & other reads without the official Twitter API.

Use an Xquik SDK for application code. Use this provider for related Twitter automation.

## Install

This provider requires Terraform CLI 1.0 or newer.

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source  = "Xquik-dev/x-twitter-scraper"
      version = "~> 0.7"
    }
  }
}

provider "x-twitter-scraper" {}
```

Then run `terraform init`.

## Authenticate

Prefer environment variables.

```sh
export X_TWITTER_SCRAPER_API_KEY="your-api-key"
```

Bearer authentication uses `X_TWITTER_SCRAPER_BEARER_TOKEN`.

Never commit credentials or Terraform state.

## Publish a tweet

Use one stable idempotency key per intended write.

```hcl
resource "x-twitter-scraper_x_tweet" "announcement" {
  account         = "@example"
  idempotency_key = "terraform-announcement-v1"
  payload_json = jsonencode({
    text = "Published through the Xquik Terraform provider."
  })
}
```

Changing the request replaces the resource. Give the new write a new key.

## Reliability & security

`bin/check-local` checks formatting, licenses, race safety, Windows builds, coverage, reproducible builds, fuzzing, CodeQL findings & known vulnerabilities.
Maintainers publish releases with `bin/release-local`, which signs the checksums with the provider's GPG key.

See [OpenSSF evidence](OPENSSF.md) for the current criteria assessment.
Treat plans and state as sensitive. API responses can contain private data.
Report vulnerabilities through [GitHub private reporting](SECURITY.md).

## Develop locally

```sh
./scripts/bootstrap
bin/check-local
```

`bin/check-local` runs every required check. See [CONTRIBUTING.md](CONTRIBUTING.md) before changing generated code.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
