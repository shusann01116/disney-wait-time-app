variable "env" {
  description = "Environment name"
  type        = string
}

variable "lambda_source_path" {
  description = "Relative path to the lambda function source code"
  type        = string
}

variable "table_name" {
  description = "Name of the DynamoDB table"
  type        = string
}
