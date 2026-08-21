# Terraform provider for X API monitors, webhooks & Twitter automation

[![CI](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper/actions/workflows/ci.yml/badge.svg)](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper/actions/workflows/ci.yml)
[![CodeQL](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper/actions/workflows/codeql.yml/badge.svg)](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper/actions/workflows/codeql.yml)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13729/badge)](https://www.bestpractices.dev/projects/13729)

Manage Xquik monitors, signed webhooks, content workflows & approved X actions as Terraform resources.

[Quickstart](docs/guides/quickstart.md) | [Provider docs](docs/index.md) | [Examples](examples) | [REST API](https://docs.xquik.com/api-reference/overview) | [OpenAPI](https://xquik.com/openapi.json) | [Security](SECURITY.md)

## Manage durable X automation

These resources manage desired state and store durable write records:

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

- Account usage and limits: `account`
- Drafts and writing styles: `draft`, `style`
- Account and keyword monitors: `monitor`, `monitor_keyword`
- Events and giveaway draws: `event`, `draw`
- Tweet, user, and account lookups: `x_tweet`, `x_user`, `x_account`
- Durable write status: `x_write_action`
- Support requests: `support_ticket`

## Run Twitter search outside Terraform

Terraform manages durable infrastructure and write records.
Use the [Xquik REST API](https://docs.xquik.com/api-reference/overview) for Twitter search, timelines, follower exports & other reads.

Use the API when you need to:

- Search tweets without the official Twitter API.
- Return structured tweet data through an X API.
- Extract an X or Twitter timeline.
- Export X followers or following accounts.

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

Then run:

```sh
terraform init
```

## Authenticate

Prefer environment variables:

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

CI checks formatting, integrity, race safety, Windows builds, coverage & reachable vulnerabilities.
Releases use pinned actions, signed checksums & GitHub attestations.

See [OpenSSF evidence](OPENSSF.md) for the current criteria assessment.
Treat plans and state as sensitive. API responses can contain private data.
Report vulnerabilities through [GitHub private reporting](SECURITY.md).

## Develop locally

```sh
./scripts/bootstrap
./scripts/lint
./scripts/test
```

See [CONTRIBUTING.md](CONTRIBUTING.md) before changing generated code.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
