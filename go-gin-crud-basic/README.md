
```
cd go-gin-crud-basic
go mod init go-gin-crud-basic
go get -u github.com/gin-gonic/gin



kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.11.0/serving-crds.yaml

kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.11.0/serving-core.yaml

kubectl apply -l knative.dev/crd-install=true -f https://github.com/knative/net-istio/releases/download/knative-v1.11.0/istio.yaml
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.11.0/istio.yaml
kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.11.0/net-istio.yaml
kind create cluster --name=my-cluster --image kindest/node:v1.27.0


Get All:
curl http://localhost:8080/users

Get By Id:
curl http://localhost:8080/users/1

Add:
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","age":28}' http://localhost:8080/users

Update:
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Bob","age":35}' http://localhost:8080/users/1

Delete:
curl -X DELETE http://localhost:8080/users/1

```

### Adding swagger documentation

To add Swagger documentation for a simple CRUD API using Gin Gonic, you can use a package called swaggo. swaggo generates Swagger documentation for your Go code and integrates with Gin. Here are the steps to add Swagger documentation to the CRUD API created using Gin Gonic for managing users:

Install dependencies and add swagger annotation to your code.
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

Install swag:
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin

swag init

2023/11/23 07:38:34 Generate swagger docs....
2023/11/23 07:38:34 Generate general API Info, search dir:./
2023/11/23 07:38:35 Generating main.User
2023/11/23 07:38:35 create docs.go at docs/docs.go
2023/11/23 07:38:35 create swagger.json at docs/swagger.json
2023/11/23 07:38:35 create swagger.yaml at docs/swagger.yaml

http://localhost:8080/swagger/index.html
```
