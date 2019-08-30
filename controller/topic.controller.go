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

func TopicRouter(router *gin.RouterGroup) {
	topicRouter := router.Group("/topic").Use(middlewares.AuthJWT())
	{
		topicRouter.POST("/create", CreateTopic)
	}
	topicRouter = router.Group("/topic")
	{
		topicRouter.GET("/", GetTopic)
	}

}

func CreateTopic(c *gin.Context) {
	var v _interface.CreateTopicDto

	if err := c.Bind(&v); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, errMes))
		return
	}

	userModel := c.MustGet("user_model").(models.User)

	if _, err := models.CreateTopic(userModel, v); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true})
		return
	}

}

func GetTopic(c *gin.Context) {
	var v _interface.GetAllTopicDto
	if err := c.Bind(&v); err != nil {
		c.JSON(http.StatusBadRequest, serializers.ErrorBaseResponse(http.StatusBadRequest, "topic_id is required"))
		return
	}
	var topic = &models.Topic{BaseModel: models.BaseModel{ID: uuid.FromStringOrNil(v.TopicID)}}
	if err := topic.GetTopic(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	c.JSON(http.StatusOK, serializers.SuccessBaseResponse(serializers.GetFullTopicResponse(*topic)))
	return

}
