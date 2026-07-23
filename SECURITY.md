# Security Policy

## Supported Code

Security fixes target the latest stable release and the `main` branch.

Upgrade to the latest stable release before requesting backports.

## Report a Vulnerability

Use [GitHub private vulnerability reporting](https://github.com/Xquik-dev/terraform-provider-x-twitter-scraper/security/advisories/new).

Never report exploit details through public issues or discussions.

Include these details:

- Affected provider and Terraform versions
- Impact and required access
- Minimal reproduction steps
- Relevant resource or data source names
- Suggested remediation, when available

Remove API keys, bearer tokens, account identifiers, and personal data.

Email [security@xquik.com](mailto:security@xquik.com) if private reporting fails.

## Response Process

Maintainers will acknowledge reports within 3 business days.

Maintainers will validate and classify reports within 14 days.

Maintainers will prioritize critical vulnerabilities immediately.

Maintainers will coordinate disclosure timing with each reporter.

Confirmed fixes require regression tests and independent review.

Maintainers will publish advisories and fixed releases when appropriate.

## Security Boundary

The provider converts Terraform configuration into documented Xquik API requests.

It stores API response fields inside Terraform state.

Treat Terraform state and plans as sensitive.

The provider never needs browser cookies or X account passwords.

Use environment variables for API keys and bearer tokens.

Only configure trusted HTTPS API endpoints.

The provider does not secure remote state storage or Terraform runners.

## Safe Harbor

We support good-faith research that follows this policy.

Avoid privacy violations, service disruption, and unnecessary data access.

Stop testing after confirming a vulnerability.

Allow reasonable time for remediation before disclosure.

We will not pursue action against compliant good-faith research.

Xquik is an independent third-party service. Not affiliated with X Corp. "Twitter" and "X" are trademarks of X Corp.
