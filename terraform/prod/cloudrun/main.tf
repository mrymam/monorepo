resource "google_cloud_run_service" "default" {
  name     = "prod-server"
  location = "asia-northeast1"
  template {
    spec {
      containers {
        image = "asia.gcr.io/${local.project_id}/prod-server:latest"
        ports {
          container_port = 9000
        }
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}