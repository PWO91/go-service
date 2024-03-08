package users

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func UNUSED(x ...interface{}) {}

type Item struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Items []Items

func SaveInMySql(Value string) {

	var db *sql.DB

    // Capture connection properties.
    cfg := mysql.Config{
        User:   "pwo91",
        Passwd: "Arbuz!2345678",
        Net:    "tcp",
        Addr:   "mysql-service-pwo91.mysql.database.azure.com:3306",
        DBName: "storage",
		AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {

    }

    pingErr := db.Ping()
    if pingErr != nil {
       
    }
    
	result, err := db.Exec("INSERT INTO proces1 (details) VALUES (?)", Value)
	if err != nil {
      
    }

	UNUSED(result)
	
}

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "Running",
	})
}

func GetItem(c *gin.Context) {

}

func PostItem(c *gin.Context) {	
	jsonData, _ := io.ReadAll(c.Request.Body)
	myString := string(jsonData[:])
	SaveInMySql(myString)

	c.JSON(http.StatusOK, "OK")
}

func DeleteItem(c *gin.Context) {

}

func AssignRoutToEngine(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/", GetItems)
	v1.POST("/add", PostItem)
}
