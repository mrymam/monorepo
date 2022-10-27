locals {
  vars = yamldecode(file("../../setting.yaml"))

  project_id = local.vars["project_id"]
  region     = local.vars["region"]
}