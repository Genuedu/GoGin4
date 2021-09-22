package controller

import (
	"training/Gogin4/entity"
	"training/Gogin4/service"
	"training/Gogin4/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

// controller calls the function from the service
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// controller calls the function from the service
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	//binding
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	// custom validation
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil

}
