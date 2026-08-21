# Xquik Terraform provider quickstart

Manage X API monitors, signed webhooks, and durable write actions with Terraform.

Use the [REST API](https://docs.xquik.com/api-reference/overview) for Twitter search and timeline extraction.

## Install

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source  = "Xquik-dev/x-twitter-scraper"
      version = "~> 0.7"
    }
  }
}
```

Run `terraform init` inside your Terraform project.

## Authenticate

Prefer an environment variable:

```sh
export X_TWITTER_SCRAPER_API_KEY="your-api-key"
```

Bearer authentication uses `X_TWITTER_SCRAPER_BEARER_TOKEN`.

Never commit credentials, Terraform state, or local plugin configuration.

## Configure

```hcl
provider "x-twitter-scraper" {}
```

Run `terraform validate` before applying changes.

## Monitor an X account

```hcl
resource "x-twitter-scraper_monitor" "product_updates" {
  username    = "xquik"
  event_types = ["tweet.new"]
}
```

## Register a signed webhook

```hcl
resource "x-twitter-scraper_webhook" "events" {
  url         = "https://example.com/xquik/webhook"
  event_types = ["tweet.new"]
}
```

Store the returned HMAC secret in a secure secret manager.

## Publish a tweet

Give each intended write one stable idempotency key.

```hcl
resource "x-twitter-scraper_x_tweet" "announcement" {
  account         = "@example"
  idempotency_key = "terraform-announcement-v1"
  payload_json = jsonencode({
    text = "Published through the Xquik Terraform provider."
  })
}
```

Changing the request creates a new action. Give it a new key.

## More documentation

- [Provider documentation](../index.md)
- [Resource examples](../../examples)
- [REST API documentation](https://docs.xquik.com/api-reference/overview)
- [OpenAPI specification](https://xquik.com/openapi.json)
- [Security policy](../../SECURITY.md)

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
