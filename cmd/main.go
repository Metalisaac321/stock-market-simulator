package main

import (
	"log"
	"net/http"

	"github.com/Metalisaac321/stock-market-simulator/internal/util"
	"github.com/gin-gonic/gin"
)

const httpAddr = ":8080"

func main() {
	server := gin.New()
	server.GET("/", handler)
	log.Fatal(server.Run(httpAddr))
}

func handler(ctx *gin.Context) {

	ctx.String(http.StatusOK, util.Welcome("Isaac"))
}
