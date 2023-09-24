data "aws_region" "current" {}

module "collector" {
  source  = "terraform-aws-modules/lambda/aws"
  version = "~> 6.0"

  function_name = "disney-waittime-app-collector-${var.env}-${random_pet.this.id}"
  description   = "Collects wait times from the Disney API and stores them in DynamoDB"

  package_type = "Zip"
  runtime      = "provided.al2"
  handler      = "bootstrap"
  source_path = [
    {
      path = var.lambda_source_path
      commands = [
        "mkdir build",
        "GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o build/bootstrap main.go",
        ":zip build/bootstrap .",
      ]
    },
  ]

  layers = [
    "arn:aws:lambda:${data.aws_region.current.id}:901920570463:layer:aws-otel-collector-arm64-ver-0-82-0:1"
  ]

  timeout       = 3
  architectures = ["arm64"]

  environment_variables = {
    DYNAMODB_TABLENAME = var.table_name
  }

  attach_tracing_policy = true
  tracing_mode          = "Active"

  attach_policy_statements = true
  policy_statements = {
    dynamo_db = {
      effect = "Allow"
      actions = [
        "dynamodb:PutItem",
        "dynamodb:Query",
        "dynamodb:Scan",
        "dynamodb:UpdateItem",
      ]

      resources = ["*"]
    },
  }

  tags = {
    Name = "disney-waittime-app-collector"
  }
}
