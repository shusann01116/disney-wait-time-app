terraform {
  source = "../../../modules//dynamodb"
}

include "root" {
  path = find_in_parent_folders()
}
