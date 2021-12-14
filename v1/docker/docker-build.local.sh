go mod vendor
docker build --tag vallinplasencia/demo-gin-rest-api-library:dev-xxx -f dockerfile.local ../.
rm -r ../vendor