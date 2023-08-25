module "collector" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "~> 6.0"

  function_name = "disney-waittime-app-collector-${var.env}"
  description   = "My awesome lambda function"
  handler       = "index.handler"
  runtime       = "python3.11"
  publish       = true
  timeout       = 60
  architectures = ["arm64"]

  source_path = var.lambda_source_path

  environment_variables = {
    DYNAMODB_TABLENAME = var.table_name
  }

  allowed_triggers = {
    CronRule = {
      principal  = "events.amazonaws.com"
      source_arn = aws_cloudwatch_event_rule.cron_rule.arn
    }
  }

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

resource "aws_cloudwatch_event_rule" "cron_rule" {
  name        = "disney-waittime-app-collector-scheduler-${var.env}"
  description = "Schedule to run the collector lambda function"

  schedule_expression = "rate(5 minutes)"
}

resource "aws_cloudwatch_event_target" "invoke_collector" {
  rule = aws_cloudwatch_event_rule.cron_rule.name
  arn  = module.collector.lambda_function_arn
}
