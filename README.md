# Terraform AWS GitHub Actions /w OIDC

This module configures AWS OIDC authentication with GitHub Actions, eliminating the need for static AWS IAM Access Keys
when running Terraform within GitHub Actions.

## GitHub Actions temporary AWS role credentials

We can use OpenID Connect within our workflows to authenticate with Amazon Web Services and get temporary session tokens
to assume IAM roles.

```mermaid
sequenceDiagram
    autonumber
    participant A as Github Actions
    participant B as Oidc Provider
    participant C as AWS IAM

    A->>+B: Request JWT
    B->>-A: Issue signed JWT
    A->>+C: Request Access Token
    C-->>+B: Verify Token
    B-->>-C: Valid
    C->>-A: Issue Role Access Session Token
```

### Usages

The module is at the time of writing super opinionated in how the AWS IAM Open ID Connect Provider is created. The only
thing the user needs to be concerned about is the `var.conditions`. This needs to be set up the trust policy for
the `sub` field, which is explained
further [here.](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services#adding-the-identity-provider-to-aws)

Multiple fields apart from the required `sub` field are supported for more granular permissions, a reference to the contents of the OIDC token can be found [here.](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#understanding-the-oidc-token)

### Examples

Check out the [examples/simple](./examples/simple/) directory for using the module.

<!-- BEGIN_TF_DOCS -->
## Resources

| Name | Type |
|------|------|
| [aws_iam_openid_connect_provider.github_actions](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_openid_connect_provider) | resource |
| [aws_iam_role.github_actions](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.github_actions](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_policy_document.github_actions](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_conditions"></a> [conditions](#input\_conditions) | Github conditions to apply to the AWS Role. E.g. from which org/repo/branch is it allowed to be run. Key is used as the JWT claim and value as the claim value. | `map(string)` | n/a | yes |
| <a name="input_policy_arn"></a> [policy\_arn](#input\_policy\_arn) | List of ARNs of IAM policies to attach to IAM role. | `list(string)` | n/a | yes |
| <a name="input_role_name"></a> [role\_name](#input\_role\_name) | The name of the AWS Role which will be used to run Github Actions. | `string` | n/a | yes |
| <a name="input_role_max_sessions_duration"></a> [role\_max\_sessions\_duration](#input\_role\_max\_sessions\_duration) | Maximum session duration (in seconds) that you want to set for the specified role. | `number` | `3600` | no |
| <a name="input_role_permission_boundary"></a> [role\_permission\_boundary](#input\_role\_permission\_boundary) | Boundary for the created role. | `string` | `null` | no |
## Outputs

| Name | Description |
|------|-------------|
| <a name="output_openid_connect_provider"></a> [openid\_connect\_provider](#output\_openid\_connect\_provider) | AWS OpenID Connected identity provider. |
| <a name="output_role"></a> [role](#output\_role) | AWS Role created |
<!-- END_TF_DOCS -->
