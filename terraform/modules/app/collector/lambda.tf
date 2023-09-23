data "aws_region" "current" {}

module "collector" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "~> 6.0"

  function_name = "disney-waittime-app-collector-${var.env}"
  description   = "My awesome lambda function"
  handler       = "index.handler"
  runtime       = "python3.11"
  publish       = true
  timeout       = 15
  architectures = ["arm64"]

  source_path = var.lambda_source_path

  environment_variables = {
    DYNAMODB_TABLENAME      = var.table_name
    AWS_LAMBDA_EXEC_WRAPPER = "/opt/otel-instrument"
  }

  layers = [
    "arn:aws:lambda:${data.aws_region.current.name}:901920570463:layer:aws-otel-python-arm64-ver-1-19-0:2"
  ]

  attach_tracing_policy = true
  tracing_mode          = "Active"

  attach_policy_statements = true
  policy_statements = {
    dynamo_db = {
      effect    = "Allow"
      actions   = ["dynamodb:PutItem"]
      resources = ["*"]
    },
  }

  tags = {
    Name = "disney-waittime-app-collector"
  }
}

resource "aws_iam_policy" "cron_rule" {
  name        = "disney-waittime-app-collector-scheduler-${var.env}"
  description = "Policy to allow the scheduler to invoke the collector lambda function"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = "lambda:InvokeFunction"
        Effect   = "Allow"
        Resource = module.collector.lambda_function_arn
      }
    ]
  })
}

resource "aws_iam_role" "cron_rule" {
  name = "disney-waittime-app-collector-scheduler-${var.env}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "scheduler.amazonaws.com"
        }
      }
    ]
  })

  managed_policy_arns = [
    aws_iam_policy.cron_rule.arn
  ]
}

resource "aws_scheduler_schedule" "cron_rule" {
  name        = "disney-waittime-app-collector-scheduler-${var.env}"
  group_name  = "default"
  description = "Schedule to run the collector lambda function"

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression = "rate(5 minutes)"

  target {
    arn      = module.collector.lambda_function_arn
    role_arn = aws_iam_role.cron_rule.arn
  }
}
