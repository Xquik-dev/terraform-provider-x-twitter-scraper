# Xquik Terraform Provider for X (Twitter) Scraper API

[![Ask DeepWiki](https://deepwiki.com/badge.svg?url=https%3A%2F%2Fgithub.com%2FXquik-dev%2Fterraform-provider-x-twitter-scraper)](https://deepwiki.com/Xquik-dev/terraform-provider-x-twitter-scraper)

Xquik Terraform Provider for the X (Twitter) Scraper API manages X data and automation workflows as infrastructure: monitors, HMAC webhooks, API keys, compose and draft resources, tweet actions, user data, draws, and extraction jobs.

Use it when Terraform should own Xquik API resources for social media automation, account monitoring, webhook delivery, and repeatable X data workflows. Start with the [Terraform quickstart](docs/guides/quickstart.md), generated [provider docs](docs/index.md), [examples](examples), or the [REST API docs](https://docs.xquik.com/api-reference/overview).

[Terraform Quickstart](docs/guides/quickstart.md) | [Provider Docs](docs/index.md) | [Examples](examples) | [REST API Docs](https://docs.xquik.com/api-reference/overview) | [OpenAPI Spec](https://xquik.com/openapi.json) | [Webhooks](https://docs.xquik.com/api-reference/webhooks/create)

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Terraform Registry publication is pending. Until the provider is listed there,
build it locally and use a Terraform CLI development override.

Build the provider binary:

```sh
./scripts/build
```

Move the generated binary into a local plugin directory:

```sh
mkdir -p ~/.terraform.d/plugins/xquik-dev/x-twitter-scraper/0.2.0/darwin_arm64
mv terraform-provider-x-twitter-scraper ~/.terraform.d/plugins/xquik-dev/x-twitter-scraper/0.2.0/darwin_arm64/
```

Use the matching platform directory for your system, such as `linux_amd64`,
`linux_arm64`, `darwin_amd64`, `darwin_arm64`, or `windows_amd64`.

Add this development override to `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "Xquik-dev/x-twitter-scraper" = "/Users/you/.terraform.d/plugins/xquik-dev/x-twitter-scraper/0.2.0/darwin_arm64"
  }

  direct {}
}
```

Then add the provider to your Terraform project:

<!-- x-release-please-start-version -->

```hcl
terraform {
  required_providers {
    x-twitter-scraper = {
      source = "Xquik-dev/x-twitter-scraper"
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
refer to the generated provider documentation in [./docs](./docs).

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
