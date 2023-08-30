resource "aws_dynamodb_table" "wating_times" {
  name         = var.table_name
  hash_key     = "attraction_id"
  range_key    = "timestamp"
  billing_mode = "PAY_PER_REQUEST"
  table_class  = "STANDARD_INFREQUENT_ACCESS"

  attribute {
    name = "attraction_id"
    type = "S"
  }
  attribute {
    name = "timestamp"
    type = "S"
  }
  attribute {
    name = "date"
    type = "S"
  }

  global_secondary_index {
    name            = "date-index"
    hash_key        = "date"
    range_key       = "attraction_id"
    projection_type = "ALL"
  }
}
