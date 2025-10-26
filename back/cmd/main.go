package main


import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/gin-gonic/gin"
	"edji-back/internal/server"
	"edji-back/config"
)
func main() {
	cfg := config.Load()
	log.Println("Server Starting")
	r := gin.Default()
  
	s := server.New(r, cfg)

	log.Println("Server Running")
	go s.Start()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SI)
}
