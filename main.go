package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", test)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("can not start server " + err.Error())
	}
}
func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "test")
}
