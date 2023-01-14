output "openid_connect_provider" {
  description = "AWS OpenID Connected identity provider."
  value       = aws_iam_openid_connect_provider.github_actions
}

output "role" {
  description = "AWS Role created"
  value       = aws_iam_role.github_actions
}
