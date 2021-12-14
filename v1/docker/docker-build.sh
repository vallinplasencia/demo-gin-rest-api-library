export SHORT_COMMIT=$(git log -1 --pretty="%H" | cut -b -8)
export DOCKER_IMAGE_VERSION="dev-${SHORT_COMMIT}"
# export GITHUB_TOKEN="token-github-access-repos-privados-golang"
docker build -t vallinplasencia/demo-gin-rest-api-library:${DOCKER_IMAGE_VERSION} --build-arg GITHUB_TOKEN -f dockerfile ../.
docker tag vallinplasencia/demo-gin-rest-api-library:${DOCKER_IMAGE_VERSION} vallinplasencia/demo-gin-rest-api-library:latest
docker push vallinplasencia/demo-gin-rest-api-library:${DOCKER_IMAGE_VERSION}
docker push vallinplasencia/demo-gin-rest-api-library:latest