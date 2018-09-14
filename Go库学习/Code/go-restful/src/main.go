package main

import (
	"log"
	"net/http"
	"userservice"

	"github.com/emicklei/go-restful"
)

func main() {
	restful.Add(userservice.New())
	log.Fatal(http.ListenAndServe(":9090", nil))
}
