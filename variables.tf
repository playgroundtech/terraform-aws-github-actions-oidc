variable "condition" {
  description = "Github conditions to apply to the AWS Role. E.g. from which org/repo/branch is it allowed to be run."
  type        = string
}

variable "policy_arn" {
  description = "List of ARNs of IAM policies to attach to IAM role."
  type        = list(string)
}

variable "role_max_sessions_duration" {
  description = "Maximum session duration (in seconds) that you want to set for the specified role."
  type        = number
  validation {
    condition     = var.role_max_sessions_duration >= 3600 && var.role_max_sessions_duration <= 43200
    error_message = "Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 3600 seconds to 43200 seconds."
  }
  default = 3600
}

variable "role_name" {
  description = "The name of the AWS Role which will be used to run Github Actions."
  type        = string
  validation {
    condition     = can(regex("^[\\w+=,.@-]{1,64}$", var.role_name))
    error_message = "Role name is invalid, only alphanumeric characters, the special characters: +=,.@- are allowed and it should be between 1 and 64 characters long."
  }
}

variable "role_permission_boundary" {
  description = "Boundary for the created role."
  type        = string
  default     = null
}
