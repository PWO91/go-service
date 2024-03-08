package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Items []Items

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func GetItem(c *gin.Context) {

}

func PostItem(c *gin.Context) {

}

func DeleteItem(c *gin.Context) {

}

func AssignRoutToEngine(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/", GetItems)
}
