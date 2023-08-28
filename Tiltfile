# Build docker images
docker_build(dockerfile='core/Dockerfile', ref='core-server', context='./core/')
docker_build(dockerfile='db/Dockerfile', ref='seed-pg', context='./')
docker_build(dockerfile='marketplace/Dockerfile', ref='next-js-frontend', context='./marketplace')

# Load k8s yaml configurations
k8s_yaml('k8s/frontend-deployment.yaml')
k8s_yaml('k8s/backend-deployment.yaml')
k8s_yaml('k8s/db-deployment.yaml')
k8s_yaml('k8s/seed.yaml')

# Set up resources 
# Note: port forwards are for localhost only, we set up k8s services for cross-pod communication
k8s_resource('core-server', port_forwards=8080)
k8s_resource('my-frontend', port_forwards=3001)
k8s_resource('postgres', port_forwards=5432)
k8s_resource('seed', resource_deps=['postgres'])

