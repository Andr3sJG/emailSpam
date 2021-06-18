package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"spam.com/controller"
	"spam.com/service"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello there <h1>")

}

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(),gin.Logger())
	

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})
	server.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})
	server.Run(":8080")

	/*port := os.Getenv("PORT")
	http.HandleFunc("/", handlerFunc)
	log.Print(port)
	http.ListenAndServe(":"+port, nil) //works on heroku*/

}
