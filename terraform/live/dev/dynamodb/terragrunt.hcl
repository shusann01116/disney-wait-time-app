terraform {
  source = "../../../modules/app//dynamodb"
}

include "root" {
  path = find_in_parent_folders()
}
