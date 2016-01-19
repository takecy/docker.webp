package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/takecy/docker.webp/img"
)

var (
	port = flag.Int("port", "3001", "")
)

func main() {
	flag.Parse()

	router := gin.Default()

	img.New(router)

	router.Run(fmt.Sprintf(":%d", *port))
}
