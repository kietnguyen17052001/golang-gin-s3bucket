package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8MiB

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			// "title": "Main website",
		})
	})

	r.POST("/upload", func(ctx *gin.Context) {
		// Get file
		file, err := ctx.FormFile("image")
		log.Println("File name:", file.Filename)
		if err != nil {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}

		// Save file
		err = ctx.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

		if err != nil {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"image": "/assets/uploads/" + file.Filename,
		})
	})

	r.Run()
}
