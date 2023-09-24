resource "aws_iam_policy" "cron_rule" {
  name        = "disney-waittime-app-collector-scheduler-${var.env}-${random_pet.this.id}"
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
  name = "disney-waittime-app-collector-scheduler-${var.env}-${random_pet.this.id}"

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
  name        = "disney-waittime-app-collector-scheduler-${var.env}-${random_pet.this.id}"
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
