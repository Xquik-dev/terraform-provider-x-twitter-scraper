# Security policy

## Supported code

Security fixes target the latest stable release and the `main` branch.

Upgrade to the latest stable release before requesting backports.

## Report a vulnerability

Use [GitHub private vulnerability reporting](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper/security/advisories/new).

Never publish exploit details in an issue or discussion.

Include:

- Affected provider and Terraform versions
- Impact and required access
- Minimal reproduction steps
- Relevant resource or data source names
- Suggested fix, when available

Remove API keys, bearer tokens, account identifiers, and personal data.

Email [security@xquik.com](mailto:security@xquik.com) if private reporting fails.

## Response process

Expect an acknowledgment within 3 business days and classification within 14
days. Maintainers prioritize critical vulnerabilities and coordinate disclosure.

Confirmed fixes require regression tests and independent review.

Maintainers will publish advisories and fixed releases when appropriate.

## Security boundary

The provider converts Terraform configuration into documented Xquik API
requests and stores response fields in Terraform state.

Treat Terraform state and plans as sensitive.

The provider never needs browser cookies or X account passwords.

Use environment variables for API keys and bearer tokens.

Only configure trusted HTTPS API endpoints.

The provider does not secure remote state storage or Terraform runners.

## Safe harbor

We support good-faith research that follows this policy.

Avoid privacy violations, service disruption, and unnecessary data access.

Stop testing after confirming a vulnerability.

Allow reasonable time for remediation before disclosure.

We will not pursue action against research that follows this policy.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
