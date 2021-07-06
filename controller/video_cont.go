package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"spam.com/enty"
	"spam.com/service"
)

type VideoController interface {
	FindAll() []enty.Video
	Save(ctx *gin.Context) enty.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []enty.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) enty.Video {
	var video enty.Video
	ctx.ShouldBindJSON(&video)
	c.service.Save(video)
	return video
}
func (c *controller)ShowAll(ctx *gin.Context){
	videos := c.service.FindAll()
	data := gin.H{
		"title" : "Video Page",
		"videos" : videos,
	}
	ctx.HTML(http.StatusOK,"main.html",data)
}
