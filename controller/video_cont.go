package controller

import (	
	"github.com/gin-gonic/gin"
	"spam.com/enty"
	"spam.com/service"
)

type VideoController interface {
	FindAll() []enty.Video
	Save(ctx *gin.Context) enty.Video
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
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
