# X (Twitter) Scraper Terraform Provider: Monitors, Webhooks, Tweet Actions & Automation

> **Xquik is an independent third-party service.** Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.

[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/13729/badge)](https://www.bestpractices.dev/projects/13729)
[![Ask DeepWiki](https://deepwiki.com/badge.svg?url=https%3A%2F%2Fgithub.com%2FXquik-dev%2Fterraform-provider-x-twitter-scraper)](https://deepwiki.com/Xquik-dev/terraform-provider-x-twitter-scraper)
[![Skills.sh x-twitter-scraper Skill](https://skills.sh/b/xquik-dev/x-twitter-scraper)](https://skills.sh/xquik-dev/x-twitter-scraper)
<a href="https://nothumansearch.ai/site/xquik.com" target="_blank" rel="noopener"><img src="https://nothumansearch.ai/badge/xquik.com.svg" alt="NHS Agentic Readiness Score" height="28"></a>

Xquik Terraform Provider for the X (Twitter) Scraper API manages Twitter automation and X data workflows as infrastructure: monitors, HMAC webhooks, API keys, compose and draft resources, tweet actions, user data, giveaway draws, follower exports, and extraction jobs.

Use it when Terraform should own Xquik API resources for social media automation, account monitoring, webhook delivery, tweet search workflows, profile tweet collection, follower export jobs, and repeatable posting infrastructure. Start with the [Terraform quickstart](docs/guides/quickstart.md), [Context7 guide](docs/CONTEXT7.md), generated [provider docs](docs/index.md), [examples](examples), or the [REST API docs](https://docs.xquik.com/api-reference/overview).

[Terraform Quickstart](docs/guides/quickstart.md) | [Context7 Guide](docs/CONTEXT7.md) | [Provider Docs](docs/index.md) | [Examples](examples) | [REST API Docs](https://docs.xquik.com/api-reference/overview) | [OpenAPI Spec](https://xquik.com/openapi.json) | [Webhooks](https://docs.xquik.com/api-reference/webhooks/create)

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on HashiCorp's website.

## Usage

Add the provider to your Terraform project:

<!-- x-release-please-start-version -->

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source  = "Xquik-dev/x-twitter-scraper"
      version = "~> 0.3.0"
    }
  }
}

provider "x-twitter-scraper" {
  # Prefer X_TWITTER_SCRAPER_API_KEY or X_TWITTER_SCRAPER_BEARER_TOKEN.
}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/Xquik-dev/x-twitter-scraper/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property     | Environment variable             | Required | Default value |
| ------------ | -------------------------------- | -------- | ------------- |
| bearer_token | `X_TWITTER_SCRAPER_BEARER_TOKEN` | false    | None          |
| api_key      | `X_TWITTER_SCRAPER_API_KEY`      | false    | None          |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/Xquik-dev/terraform-provider-x-twitter-scraper/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
