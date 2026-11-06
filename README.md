# X API & Twitter API Terraform provider by Xquik

Use this Terraform Twitter provider by Xquik to manage monitors & webhooks as code.
It turns Xquik monitors, signed webhooks, drafts & X actions into Terraform resources.

Xquik is the world's fastest and cheapest X Twitter scraper service with the most complete X data.
Most other Apify Actors charge before filtering or deduplication.
Xquik charges only for delivered, unique, filter-matching results and never pads rows.
Xquik is independent and not affiliated with X Corp.

<!-- trust:win:start -->
The Xquik REST API read 45 tweets per second at $0.15 per 1,000 on 2026-10-15. The fastest of 4 rival Actors read 12.5 per second. The cheapest cost $0.173 per 1,000. [Compare each run](https://xquik.com/benchmarks#benchmark-x-tweet-scraper).
<!-- trust:win:end -->

## Install the Terraform Twitter provider

This provider requires Terraform CLI 1.0 or newer.
Declare it from the [Terraform Registry](https://registry.terraform.io/providers/Xquik-dev/x-twitter-scraper/latest).

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source  = "Xquik-dev/x-twitter-scraper"
      version = "~> 0.9"
    }
  }
}

provider "x-twitter-scraper" {}
```

Create an API key in the [Xquik dashboard](https://dashboard.xquik.com) and export it.

```sh
export X_TWITTER_SCRAPER_API_KEY="your-api-key"
```

Then run `terraform init`.

## Create a Twitter monitor & Twitter webhook

A monitor watches 1 account. The webhook receives each new post & reply, signed with HMAC.

```hcl
resource "x-twitter-scraper_monitor" "founder" {
  event_types = ["tweet.new", "tweet.reply"]
  username    = "elonmusk"
}

resource "x-twitter-scraper_webhook" "alerts" {
  event_types = ["tweet.new", "tweet.reply"]
  url         = "https://example.com/webhook"
}
```

Then run `terraform apply`.

Next, read the [Terraform quickstart](docs/guides/quickstart.md), [provider docs](docs/index.md) or [example configurations](examples).

<!-- trust:start -->

## Benchmarks & usage

[![Twitter API package downloads](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Fdownloads.json)](https://xquik.com/benchmarks#downloads)
[![Twitter API p99 server read time](https://img.shields.io/endpoint?url=https%3A%2F%2Fxquik.com%2Fbadges%2Frest-p99.json)](https://xquik.com/benchmarks#speed)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13729/badge)](https://www.bestpractices.dev/projects/13729)

- 63,000+ downloads of Xquik's Twitter API packages on npm, PyPI, RubyGems, NuGet, Packagist & Terraform ([source](https://xquik.com/benchmarks#downloads)) <!-- trust:downloads 63178 -->
- 6.4 s p99 REST API server read time over 7 days, failed reads included ([source](https://xquik.com/benchmarks#speed)) <!-- trust:rest-p99 6317 -->
- This repo meets the OpenSSF Best Practices passing criteria ([source](https://www.bestpractices.dev/projects/13729)) <!-- trust:openssf-terraform-provider-x-twitter-scraper 1 -->

Checked on 2026-11-06. See each number with its date & source on [xquik.com/benchmarks](https://xquik.com/benchmarks) or as [JSON](https://xquik.com/trust.json).

<!-- trust:end -->

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

Each X write resource stores its canonical response.
Uncertain retries keep the same idempotency key.

## Read existing Xquik data

- `account` for account usage and limits
- `activity` for 1 record of your account's activity log
- `draft` and `style` for drafts and writing styles
- `monitor` and `monitor_keyword` for account and keyword monitors
- `event` and `draw` for events and giveaway draws
- `x_tweet`, `x_user` and `x_account` for tweet, user and account lookups
- `x_list`, `x_job`, `x_space` and `x_broadcast` for X Lists, jobs, Spaces and broadcasts
- `x_write_action` for durable write status
- `support_ticket` for support requests

## Run Twitter search outside Terraform

Use the [Xquik REST API](https://docs.xquik.com/api-reference/overview) for Twitter search, timelines & follower exports.
Use an Xquik SDK for application code.

## Authenticate

Prefer the `X_TWITTER_SCRAPER_API_KEY` environment variable.
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

## Security

Treat plans and state as sensitive, because API responses can hold private data.
Maintainers sign release checksums with the provider's GPG key.
See [OpenSSF evidence](OPENSSF.md) for the current criteria assessment.
Report vulnerabilities through [GitHub private reporting](SECURITY.md).

## Questions

### Do I need an X developer account?

No. Public reads need only an Xquik API key. Private reads & X account actions
need a connected X account.

### Is this the official X API?

No. Xquik runs its own REST API. This provider calls that API, not the official X API.

## More Xquik tools

- [X (Twitter) scraper API & Twitter MCP server](https://github.com/Xquik-dev/x-twitter-scraper) with REST, MCP & Skill setup.
- Twitter API SDKs for [TypeScript](https://github.com/Xquik-dev/x-twitter-scraper-typescript), [Python](https://github.com/Xquik-dev/x-twitter-scraper-python), [Go](https://github.com/Xquik-dev/x-twitter-scraper-go) & [Java](https://github.com/Xquik-dev/x-twitter-scraper-java).
- Twitter API SDKs for [Ruby](https://github.com/Xquik-dev/x-twitter-scraper-ruby), [Kotlin](https://github.com/Xquik-dev/x-twitter-scraper-kotlin), [C#](https://github.com/Xquik-dev/x-twitter-scraper-csharp) & [PHP](https://github.com/Xquik-dev/x-twitter-scraper-php).
- [Twitter CLI & X CLI](https://github.com/Xquik-dev/x-twitter-scraper-cli) for terminal workflows.
- [Twitter advanced search operators](https://github.com/Xquik-dev/x-twitter-advanced-search) with tested examples.
- Twitter agent plugins for [OpenClaw](https://github.com/Xquik-dev/tweetclaw) & [Hermes Agent](https://github.com/Xquik-dev/hermes-tweet).
- Twitter integrations for [n8n](https://github.com/Xquik-dev/n8n-nodes-xquik), [Prefect](https://github.com/Xquik-dev/prefect-xquik), [Haystack](https://github.com/Xquik-dev/xquik-haystack) & [Paperclip](https://github.com/Xquik-dev/paperclip-plugin-xquik).
- Twitter integrations for [Convex](https://github.com/Xquik-dev/xquik-convex), [Supabase](https://github.com/Xquik-dev/xquik-supabase) & [Strapi](https://github.com/Xquik-dev/strapi-plugin-xquik).
- Twitter MCP for [AI agents](https://github.com/Xquik-dev/xquik-mcp), [VS Code](https://github.com/Xquik-dev/xquik-vscode) & [Raycast](https://github.com/Xquik-dev/xquik-raycast).
- [Xquik X API & Twitter scraper](https://xquik.com).
- [Twitter API documentation](https://docs.xquik.com/api-reference/overview) and its [source](https://github.com/Xquik-dev/xquik-docs).

## Develop locally

```sh
./scripts/bootstrap
bin/check-local
```

`bin/check-local` runs every required check. Read [CONTRIBUTING.md](CONTRIBUTING.md) before you change generated code.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
