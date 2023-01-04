env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
docker build -t devops_test . --no-cache
# docker push {registry_domain}

