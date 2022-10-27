# monorepo


## terraform

in terraform directory

```bash
terraform apply 
```

### gcr repository

create GCR repository before terraform apply cloudrun resource
```bash
export PROJECT_ID=<project_id>
gcloud auth login
gcloud auth configure-docker
docker build -f server/docker/Dockerfile -t asia.gcr.io/$PROJECT_ID/prod-server:latest .
docker push asia.gcr.io/$PROJECT_ID/prod-server:latest
```