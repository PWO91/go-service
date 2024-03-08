package main

import (
	"template/service/handlers/users"

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

    router.Run("localhost:8080")

}