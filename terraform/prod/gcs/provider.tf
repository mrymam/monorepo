provider "google" {
  project = local.project_id
  region  = local.region
}

terraform {
  backend "gcs" {
    prefix = "terraform/prod/gcs"
  }
}
