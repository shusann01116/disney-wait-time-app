resource "aws_s3_bucket" "tdl_collector_bucket" {
  bucket = "tdr-collector-bucket-${var.env}"
}
