# go initialize
```sh
cd go-gin-hello
go mod init go-gin-hello
go get -u github.com/gin-gonic/gin
```
# run local
```sh
export MESSAGE="New Team"
export COLOR="red"
go run cmd/hello/main.go
curl http://localhost:8080
```

# build docker image
```sh
docker build -t go-gin-hello .
docker run -d -p 8080:8080 go-gin-hello

docker tag go-gin-hello:latest rsxyz123/go-gin-hello:latest
docker push rsxyz123/go-gin-hello:latest
docker rmi rsxyz123/go-gin-hello:latest
docker pull rsxyz123/go-gin-hello:latest
docker run -e "MESSAGE=Good Morning" -e "COLOR=blue" -p 8080:8080 rsxyz123/go-gin-hello:latest
```