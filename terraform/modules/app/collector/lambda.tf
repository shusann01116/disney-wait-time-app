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
    S3_BUCKETNAME           = aws_s3_bucket.tdl_collector_bucket.id
  }

  layers = [
    "arn:aws:lambda:${data.aws_region.current.name}:901920570463:layer:aws-otel-python-arm64-ver-1-19-0:2"
  ]

  allowed_triggers = {
    CronRule = {
      principal  = "events.amazonaws.com"
      source_arn = aws_cloudwatch_event_rule.cron_rule.arn
    }
  }

  attach_tracing_policy = true
  tracing_mode          = "Active"

  attach_policy_statements = true
  policy_statements = {
    dynamo_db = {
      effect    = "Allow"
      actions   = ["dynamodb:PutItem"]
      resources = ["*"]
    },
    s3 = {
      effect    = "Allow"
      actions   = ["s3:PutObject"]
      resources = ["${aws_s3_bucket.tdl_collector_bucket.arn}/*"]
    }
  }

  tags = {
    Name = "disney-waittime-app-collector"
  }
}

resource "aws_cloudwatch_event_rule" "cron_rule" {
  name        = "disney-waittime-app-collector-scheduler-${var.env}"
  description = "Schedule to run the collector lambda function"

  schedule_expression = "rate(5 minutes)"
}

resource "aws_cloudwatch_event_target" "invoke_collector" {
  rule = aws_cloudwatch_event_rule.cron_rule.name
  arn  = module.collector.lambda_function_arn
}
