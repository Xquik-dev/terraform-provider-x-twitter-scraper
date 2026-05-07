# Xquik Terraform Provider Quickstart

Use the Xquik Terraform Provider to manage X (Twitter) automation resources from Terraform. The provider covers Xquik API objects such as monitors, HMAC webhooks, API keys, drafts, compose workflows, tweet actions, user data, draws, and extraction jobs.

## Install

Terraform Registry publication is pending. Until the provider is listed there, build the provider locally and use a Terraform CLI development override.

```sh
./scripts/build
mkdir -p ~/.terraform.d/plugins/xquik-dev/x-twitter-scraper/0.2.0/darwin_arm64
mv terraform-provider-x-twitter-scraper ~/.terraform.d/plugins/xquik-dev/x-twitter-scraper/0.2.0/darwin_arm64/
```

Use the platform directory that matches your machine, such as `linux_amd64`, `linux_arm64`, `darwin_amd64`, `darwin_arm64`, or `windows_amd64`.

Add a local override to `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "Xquik-dev/x-twitter-scraper" = "/Users/you/.terraform.d/plugins/xquik-dev/x-twitter-scraper/0.2.0/darwin_arm64"
  }

  direct {}
}
```

## Authenticate

Prefer environment variables for credentials:

```sh
export X_TWITTER_SCRAPER_API_KEY="xqk_your_key"
```

The provider also supports `X_TWITTER_SCRAPER_BEARER_TOKEN` for bearer token authentication. Do not commit API keys, bearer tokens, Terraform state files, local plugin binaries, or machine-specific `.terraformrc` paths.

## Configure

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source = "Xquik-dev/x-twitter-scraper"
    }
  }
}

provider "x-twitter-scraper" {}
```

Run `terraform init` from your Terraform project directory after the local override is configured.

## First Resources

Create a monitor for new account activity:

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

## Documentation

- Generated provider docs: [docs/index.md](../index.md)
- Resource examples: [examples](../../examples)
- REST API docs: https://docs.xquik.com/api-reference/overview
- Webhook docs: https://docs.xquik.com/api-reference/webhooks/create
- OpenAPI spec: https://xquik.com/openapi.json
