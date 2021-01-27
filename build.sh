CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api/oauth2-test
docker build -t 192.168.247.22/ima/oauth2-test .
docker rmi $(docker images -f "dangling=true" -q) -f || true
docker push 192.168.247.22/ima/oauth2-test