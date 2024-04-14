Framework: gin

go mod init

go get github.com/gin-gonic/gin

curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"