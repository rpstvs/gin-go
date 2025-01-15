package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rpstvs/gin-go/internal/api"
)

func main() {
	r := gin.Default()
	token, err := api.GetToken(,)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"token": token,
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "helloword")
	})
	r.Run()
}
