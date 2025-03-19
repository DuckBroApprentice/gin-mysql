package main

import (
	"tryweb/database"
	"tryweb/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	user := r.Group("/user")
	router.UserRouter(user)

	go func() {
		database.DBConn()
	}()

	r.Run()
}
