package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"spam.com/controller"
	"spam.com/repository"
	"spam.com/service"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Hello there <h1>")

}

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	server.Static("/css", "./css")
	server.LoadHTMLGlob("html/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/test", func(ctx *gin.Context) {

			ctx.JSON(200, videoController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	
	port := os.Getenv("PORT")
	server.Run(":" + port)
	
	//server.Run(":8080")
	/*port := os.Getenv("PORT")
	http.HandleFunc("/", handlerFunc)
	log.Print(port)
	http.ListenAndServe(":"+port, nil) //works on heroku*/

}
