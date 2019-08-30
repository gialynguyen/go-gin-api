package controller

import (
	"github.com/gin-gonic/gin"
	_interface "github.com/golang-gin/interface"
	"github.com/golang-gin/middlewares"
	"github.com/golang-gin/models"
	"github.com/golang-gin/serializers"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

func TutorialRouter(router *gin.RouterGroup) {
	tutorialRouter := router.Group("/tutorial").Use(middlewares.AuthJWT())
	{
		tutorialRouter.POST("/create", CreateTutorial)
	}

	tutorialRouter = router.Group("/tutorial")
	{
		tutorialRouter.GET("/", GetTutorialById)
	}
}

func CreateTutorial (c *gin.Context) {
	var v _interface.CreateTutorialDto

	if err := c.Bind(&v); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, errMes))
		return
	}

	userModel := c.MustGet("user_model").(models.User)

	if err := models.CreateTutorial(&userModel, v); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}

func GetTutorialById(c *gin.Context) {

	v := struct {
		TutorialId string `form:"tutorial_id" json:"tutorial_id" binding:"required"`
	}{}

	if err := c.Bind(&v); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, errMes))
		return
	}

	tutorial := &models.Tutorial{BaseModel: models.BaseModel{ID: uuid.FromStringOrNil(v.TutorialId)}}
	if err := tutorial.GetAllLesson(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	resData := map[string]interface{}{
		"lesson": serializers.LessonOfTutorial(*tutorial),
	}

	c.JSON(http.StatusOK, serializers.SuccessBaseResponse(resData))
	return
}