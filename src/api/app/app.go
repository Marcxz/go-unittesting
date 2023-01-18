package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()

	if err := router.Run(":3000"); err != nil {
		fmt.Println("The server can't run on the port 3000")
		panic(3000)
	} else {
		fmt.Println("The server is running on port 3000")
	}
}
