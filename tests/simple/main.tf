data "aws_iam_policy" "view_only" {
  name = "ViewOnlyAccess"
}

module "aws_github_actions_oidc" {
  source     = "../../"
  role_name  = var.role_name
  condition  = "playgroundtech/terraform-aws-github-actions-oidc:*"
  policy_arn = [data.aws_iam_policy.view_only.arn]
}
