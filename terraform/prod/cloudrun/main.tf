resource "google_cloud_run_service" "default" {
  name     = "prod-server"
  location = "asia-northeast1"
  template {
    spec {
      containers {
        image = "asia.gcr.io/${local.project_id}/prod-server"
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}