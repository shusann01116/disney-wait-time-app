terraform {
  source = "../../../../modules/app//collector"
}

include "root" {
  path = find_in_parent_folders()
}

inputs = {
  env                = "sandbox"
  lambda_source_path = "${get_repo_root()}/app/collector_go"
}
