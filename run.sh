CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o server  
docker build -t go-docker-wercker .
docker run -p 8080:80 go-docker-wercker