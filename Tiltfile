
watch_file('assets')
# Build docker images
docker_build(dockerfile='core/Dockerfile', ref='core-server', context='./core/')
docker_build(dockerfile='db/Dockerfile', ref='seed-pg', context='./')


# Load k8s yaml configurations
k8s_yaml('k8s/backend-deployment.yaml')
k8s_yaml('k8s/db-deployment.yaml')
k8s_yaml('k8s/seed.yaml')

# Set up resources 
k8s_resource('core-server', port_forwards=8080)
k8s_resource('postgres', port_forwards=5432)
k8s_resource('seed', resource_deps=['postgres'])

