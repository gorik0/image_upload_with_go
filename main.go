package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()
	router.POST("/image", imageHandler)
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}

func imageHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error getting file: %v", err)
	}
	println(file.Filename)
}
