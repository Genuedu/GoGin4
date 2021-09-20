package controller

import (
	"training/Gogin4/entity"
	"training/Gogin4/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

// controller calls the function from the service
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// controller calls the function from the service
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video

}
