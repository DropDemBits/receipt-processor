package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		data := map[string]any{
			"hello": "world",
		}

		ctx.JSON(http.StatusOK, data)
	})

	r.Run()
}
