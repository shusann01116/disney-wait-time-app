data "aws_region" "current" {}

module "appsync" {
  source  = "terraform-aws-modules/appsync/aws"
  version = "2.2.0"

  name = "disney-waittime-app-appsync-${var.env}"

  schema = file("${path.module}/src/schema.graphql")

  visibility = "GLOBAL"

  api_keys = {
    default = null # such key will expire in 7 days
  }

  datasources = {
    waittime_table = {
      type       = "AMAZON_DYNAMODB"
      table_name = var.table_name
      region     = data.aws_region.current.name
    }
  }

  resolvers = {
    # "Query.getZip" = {
    #   data_source   = "lambda_create_zip"
    #   direct_lambda = true
    # }
    #
    # "Query.getModuleFromRegistry" = {
    #   data_source       = "registry_terraform_io"
    #   request_template  = file("vtl-templates/request.Query.getModuleFromRegistry.vtl")
    #   response_template = file("vtl-templates/response.Query.getModuleFromRegistry.vtl")
    # }
  }
}
