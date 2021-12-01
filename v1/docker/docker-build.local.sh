go mod vendor
docker build --tag vallinplasencia/demo-books:dev-xxx -f dockerfile.local ../.
rm -r ../vendor