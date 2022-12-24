package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", test)

	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal("can not start server " + err.Error())
	}
}
func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "test")
}
