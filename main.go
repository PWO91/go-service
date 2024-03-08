package main

import (
	"template/service/handlers/users"
	"os"
	"github.com/gin-gonic/gin"
)

type routes struct {
    router *gin.Engine
}

func (r routes) Run(addr ...string) error {
    return r.router.Run()
}

func main() {

	router :=gin.Default()

	users.AssignRoutToEngine(router)

	port := os.Getenv("HTTP_PLATFORM_PORT")

	if port == "" {
		port = "8080"
	}

    router.Run("127.0.0.1:"+port)

}