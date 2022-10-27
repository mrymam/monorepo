resource "google_storage_bucket" "sqlite" {
  name          = "prod-monorepo-sqlite"
  location      = "ASIA"
  force_destroy = false

  labels = {
    app = "monorepo-sqlite"
    env = "prod"
  }
}