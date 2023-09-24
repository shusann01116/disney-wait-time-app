resource "aws_dynamodb_table" "wating_times" {
  name         = var.table_name
  hash_key     = "PK"
  range_key    = "SK"
  billing_mode = "PAY_PER_REQUEST"
  table_class  = "STANDARD_INFREQUENT_ACCESS"

  attribute {
    name = "PK"
    type = "S"
  }
  attribute {
    name = "SK"
    type = "S"
  }
  attribute {
    name = "Data"
    type = "S"
  }
  attribute {
    name = "FacilityIndexId"
    type = "S"
  }

  global_secondary_index {
    name            = "GSI-1"
    hash_key        = "SK"
    range_key       = "Data"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "FacilityIndex"
    hash_key        = "FacilityIndexId"
    range_key       = "Data"
    projection_type = "INCLUDE"
    non_key_attributes = [
      "FacilityKanaName"
    ]
  }
}
