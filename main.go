package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"spam.com/controller"
	"spam.com/service"
	"os"
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
	server.Use(gin.Recovery(), gin.Logger())

	server.Static("/css", "./css")
	server.LoadHTMLGlob("html/*.html")
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})
		apiRoutes.POST("/test", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	port := os.Getenv("PORT")
	server.Run(":"+port)

	/*port := os.Getenv("PORT")
	http.HandleFunc("/", handlerFunc)
	log.Print(port)
	http.ListenAndServe(":"+port, nil) //works on heroku*/

}
