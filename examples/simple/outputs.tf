output "openid_connect_provider" {
  description = "AWS OpenID Connected identity provider."
  value       = module.aws_github_actions_oidc.openid_connect_provider
}

output "role_name" {
  description = "AWS Role created"
  value       = module.aws_github_actions_oidc.role.name
}
