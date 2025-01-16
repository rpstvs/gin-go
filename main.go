package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rpstvs/gin-go/internal/api"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("couldnt load .env file")
	}

	clientid := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")

	token, err := api.GetToken(clientid, client_secret)

	api.GetNewReleases()

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
