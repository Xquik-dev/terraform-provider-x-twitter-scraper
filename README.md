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

<!-- trust:start -->

## Trust & usage

[![Twitter API package downloads](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fdownloads.json)](https://xquik.com/trust#downloads)
[![Twitter scraper users on Apify](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fapify-users.json)](https://apify.com/xquik)
[![Twitter scraper runs on Apify](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fapify-runs.json)](https://apify.com/xquik)
[![Tweet scraper rating on Apify](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fapify-rating.json)](https://apify.com/xquik/x-tweet-scraper/reviews)
[![Xquik Twitter API accounts](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Faccounts.json)](https://xquik.com/trust#accounts)
[![Twitter Skill installs](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fskill-installs.json)](https://xquik.com/trust#skill-installs)
[![Xquik verified revenue on TrustMRR](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Frevenue.json)](https://trustmrr.com/startup/xquik)
[![Twitter API median response time](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Frest-p50.json)](https://xquik.com/trust#speed)
[![Xquik GitHub stars](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fgithub-stars.json)](https://github.com/Xquik-dev)
[![x-twitter-scraper downloads on Terraform Registry](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fterraform-x-twitter-scraper.json)](https://registry.terraform.io/providers/Xquik-dev/x-twitter-scraper/latest)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13729/badge)](https://www.bestpractices.dev/projects/13729)

- 51,000+ downloads of Xquik's Twitter API packages on npm, PyPI, RubyGems, NuGet, Packagist & Terraform ([source](https://xquik.com/trust#downloads)) <!-- trust:downloads 51020 -->
- 7,800+ users of Xquik's Twitter scraper Actors on Apify ([source](https://apify.com/xquik)) <!-- trust:apify-users 7800 -->
- 5,700,000+ runs of Xquik's Twitter scraper Actors on Apify ([source](https://apify.com/xquik)) <!-- trust:apify-runs 5704701 -->
- 4.66 of 5 for the X Tweet Scraper from 16 Apify reviews ([source](https://apify.com/xquik/x-tweet-scraper/reviews)) <!-- trust:apify-rating 4.66 16 -->
- 3,900+ accounts on xquik.com ([source](https://xquik.com/trust#accounts)) <!-- trust:accounts 3924 -->
- 1,600+ installs of Xquik's Twitter Skills on skills.sh ([source](https://xquik.com/trust#skill-installs)) <!-- trust:skill-installs 1612 -->
- $38,000+ in revenue verified by TrustMRR ([source](https://trustmrr.com/startup/xquik)) <!-- trust:revenue 38115 -->
- 501 ms median REST API read time over the last 7 days ([source](https://xquik.com/trust#speed)) <!-- trust:rest-p50 501 -->
- 370+ GitHub stars across Xquik's repos ([source](https://github.com/Xquik-dev)) <!-- trust:github-stars 374 -->
- 110+ downloads of x-twitter-scraper on Terraform Registry ([source](https://registry.terraform.io/providers/Xquik-dev/x-twitter-scraper/latest)) <!-- trust:terraform-x-twitter-scraper 115 -->
- This repo meets the OpenSSF Best Practices passing criteria ([source](https://www.bestpractices.dev/projects/13729)) <!-- trust:openssf-terraform-provider-x-twitter-scraper 1 -->

Checked on 2026-10-15. See each number with its date & source on [xquik.com/trust](https://xquik.com/trust) or as [JSON](https://xquik.com/trust.json).

<!-- trust:end -->

## Manage durable X automation

| Workflow | Resources |
| --- | --- |
| Compose and refine posts | `compose`, `draft`, `style` |
| Monitor accounts, keywords and posts | `monitor`, `monitor_keyword`, `monitor_post` |
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
- `monitor`, `monitor_keyword` and `monitor_post` for account, keyword and post monitors
- `event` and `draw` for events and giveaway draws
- `x_tweet`, `x_user` and `x_account` for tweet, user and account lookups
- `x_list`, `x_job`, `x_space` and `x_broadcast` for X Lists, job listings, Spaces and live broadcasts
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

`bin/check-local` checks API coverage, formatting, licenses, race safety, Windows builds, coverage, reproducible builds, fuzzing, CodeQL findings & known vulnerabilities.
Maintainers publish releases with `bin/release-local`, which signs the checksums with the provider's GPG key.

See [OpenSSF evidence](OPENSSF.md) for the current criteria assessment.
Treat plans and state as sensitive. API responses can contain private data.
Report vulnerabilities through [GitHub private reporting](SECURITY.md).

## More Xquik tools

- [X (Twitter) scraper API & Twitter MCP server](https://github.com/Xquik-dev/x-twitter-scraper) with REST, MCP & Skill setup.
- Twitter API SDKs for [TypeScript](https://github.com/Xquik-dev/x-twitter-scraper-typescript), [Python](https://github.com/Xquik-dev/x-twitter-scraper-python), [Go](https://github.com/Xquik-dev/x-twitter-scraper-go), [Ruby](https://github.com/Xquik-dev/x-twitter-scraper-ruby), [Java](https://github.com/Xquik-dev/x-twitter-scraper-java), [Kotlin](https://github.com/Xquik-dev/x-twitter-scraper-kotlin), [C#](https://github.com/Xquik-dev/x-twitter-scraper-csharp) & [PHP](https://github.com/Xquik-dev/x-twitter-scraper-php).
- [Twitter CLI & X CLI](https://github.com/Xquik-dev/x-twitter-scraper-cli) for terminal workflows.
- [Twitter advanced search operators](https://github.com/Xquik-dev/x-twitter-advanced-search) with tested examples.
- [Xquik X API & Twitter scraper](https://xquik.com) and the [Twitter API documentation](https://docs.xquik.com/api-reference/overview).

## Develop locally

```sh
./scripts/bootstrap
bin/check-local
```

`bin/check-local` runs every required check. See [CONTRIBUTING.md](CONTRIBUTING.md) before changing generated code.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
