locals {
  vars = yamldecode(file("../../setting.yaml"))
  tf = local.vars["terraform"]

  project _id = local.tf["project_id"]
  region      = local.tf["region"]
}