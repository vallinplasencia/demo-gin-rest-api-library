export SHORT_COMMIT=$(git log -1 --pretty="%H" | cut -b -8)
export DOCKER_IMAGE_VERSION="dev-${SHORT_COMMIT}"
# export GITHUB_TOKEN="token-github-access-repos-privados-golang"
docker build -t vallinplasencia/demo-books:${DOCKER_IMAGE_VERSION} --build-arg GITHUB_TOKEN -f dockerfile ../.
docker tag vallinplasencia/demo-books:${DOCKER_IMAGE_VERSION} vallinplasencia/demo-books:latest
docker push vallinplasencia/demo-books:${DOCKER_IMAGE_VERSION}
docker push vallinplasencia/demo-books:latest