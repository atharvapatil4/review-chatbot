docker build -t seed-img:latest -f db/Dockerfile .
kubectl apply -f k8s/seed.yaml
