package controller

import (
	"github.com/gin-gonic/gin"
	_interface "github.com/golang-gin/interface"
	"github.com/golang-gin/middlewares"
	"github.com/golang-gin/models"
	"github.com/golang-gin/serializers"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

func LessonRouter(router *gin.RouterGroup) {
	lessonRouter := router.Group("/class").Use(middlewares.AuthJWT())
	{
		lessonRouter.POST("/create", CreateClass)
	}
}

func CreateClass (c *gin.Context) {
	var v _interface.CreateLessonDto

	if err := c.Bind(&v); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, errMes))
		return
	}

	userModel := c.MustGet("user_model").(models.User)

	if err := models.CreateLesson(&userModel, v); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}

