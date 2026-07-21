# Context7 Guide

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.

This guide gives coding agents a compact path through the Xquik Terraform Provider for X (Twitter) Scraper API. Use it before generated resource pages when you need install, authentication, resource naming, and first-use examples in one place.

## Install

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source  = "Xquik-dev/x-twitter-scraper"
      version = "~> 0.2.3"
    }
  }
}
```

Run `terraform init` from your Terraform project directory.

## Authenticate

Prefer environment variables so credentials stay out of Terraform files and shell history:

```sh
export X_TWITTER_SCRAPER_API_KEY="xqk_your_key"
```

Bearer token authentication is also supported through `X_TWITTER_SCRAPER_BEARER_TOKEN`. Do not commit API keys, bearer tokens, Terraform state files, local plugin binaries, or machine-specific `.terraformrc` paths.

## Configure

```hcl
provider "x-twitter-scraper" {}
```

Validate the configuration before adding resources:

```sh
terraform providers
terraform validate
```

## First Resources

Create a monitor for account activity:

```hcl
resource "x-twitter-scraper_monitor" "product_updates" {
  username = "xquik"
  event_types = ["tweet.new"]
}
```

Register a webhook endpoint for monitor events:

```hcl
resource "x-twitter-scraper_webhook" "events" {
  url = "https://example.com/xquik/webhook"
  event_types = ["tweet.new"]
}
```

Create an API key for automation workloads:

```hcl
resource "x-twitter-scraper_api_key" "automation" {
  name = "terraform-automation"
}
```

Read the current account state with a data source:

```hcl
data "x-twitter-scraper_account" "current" {}

output "xquik_plan" {
  value = data.x-twitter-scraper_account.current.plan
}
```

## Common Resources

- `x-twitter-scraper_monitor`: monitor an X username for events.
- `x-twitter-scraper_webhook`: deliver monitor and workflow events to an HTTPS endpoint.
- `x-twitter-scraper_api_key`: create an API key for automation.
- `x-twitter-scraper_draft`: manage reusable draft content.
- `x-twitter-scraper_style`: manage tweet style presets.
- `x-twitter-scraper_x_tweet`: create or manage tweet workflows.
- `x-twitter-scraper_x_tweet_like`: manage tweet like actions.
- `x-twitter-scraper_x_tweet_retweet`: manage tweet retweet actions.
- `x-twitter-scraper_x_user_follow`: manage user follow actions.

Use [docs/resources](resources) for generated schema details and [examples/resources](../examples/resources) for copyable resource blocks.

## Data Sources

Use data sources when Terraform needs to read existing Xquik API state:

- `x-twitter-scraper_account`: account profile and plan state.
- `x-twitter-scraper_monitor`: existing monitor metadata.
- `x-twitter-scraper_webhook`: webhook endpoint metadata.
- `x-twitter-scraper_x_tweet`: tweet data.
- `x-twitter-scraper_x_user`: user profile data.
- `x-twitter-scraper_extraction`: extraction job data.
- `x-twitter-scraper_draw`: giveaway draw data.

Use [docs/data-sources](data-sources) for generated schema details and [examples/data-sources](../examples/data-sources) for data-source blocks.

## External Docs

- REST API docs: https://docs.xquik.com/api-reference/overview
- Webhooks: https://docs.xquik.com/api-reference/webhooks/create
- OpenAPI spec: https://xquik.com/openapi.json
- GitHub repo: https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper

## Agent Rules

- Use `docs/guides/quickstart.md` and this file before generated schema pages.
- Use source `Xquik-dev/x-twitter-scraper` in `required_providers`.
- Prefer `X_TWITTER_SCRAPER_API_KEY` or `X_TWITTER_SCRAPER_BEARER_TOKEN`.
- Keep credentials in environment variables or approved secret stores.
- Never commit Terraform state, local plugin binaries, `.terraformrc`, API keys, bearer tokens, or logs.
