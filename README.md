# monorepo

```bash
export PROJECT_ID=monorepo
gcloud auth login
gcloud auth configure-docker
docker build -f server/docker/Dockerfile -t asia.gcr.io/$PROJECT_ID/prod-server:latest .
docker push asia.gcr.io/$PROJECT_ID/prod-server:latest
```