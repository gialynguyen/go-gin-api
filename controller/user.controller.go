package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-gin/interface"
	"github.com/golang-gin/middlewares"
	"github.com/golang-gin/models"
	"github.com/golang-gin/serializers"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

func UserRouter(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", SignUp)
		authRouter.POST("/login", Login)
	}
	userRouter := router.Group("/users").Use(middlewares.AuthJWT())
	{
		userRouter.GET("/me", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user_name": c.MustGet("user_model").(models.User).Username,
			})
		})
		userRouter.GET("/topic", GetAllTopic)
		userRouter.GET("/profile", GetProfile)
		userRouter.POST("/follow-topic", FollowTopic)
	}

}

func SignUp(c *gin.Context) {
	var u _interface.CreateUserDto
	var err error

	if err = c.Bind(&u); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, errMes))
		return
	}

	var user models.User

	if user, err = models.CreateUser(u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, "Email has exists"))
		return
	}
	c.JSON(http.StatusCreated, serializers.SuccessBaseResponse(serializers.RegisterResponse(user)))
}

func Login(c *gin.Context) {
	var u _interface.LoginUserDto
	if err := c.Bind(&u); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)

		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, errMes))
		return
	}

	userModel, err := models.FindOne(u)

	if err != nil {
		c.JSON(http.StatusForbidden, serializers.ErrorBaseResponse(http.StatusForbidden, "Not Registered email or invalid password"))
		return
	}

	if err = userModel.CheckPassword(u.Password); err != nil {
		c.JSON(http.StatusForbidden, serializers.ErrorBaseResponse(http.StatusForbidden, "Not Registered email or invalid password"))
		return
	}

	c.Set("user_model", userModel)

	c.JSON(http.StatusOK, serializers.SuccessBaseResponse(serializers.LoginResponse(c)))
}

func GetAllTopic(c *gin.Context) {
	userModel := c.MustGet("user_model").(models.User)

	if err := userModel.GetTopic(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, "Something error"))
		return
	}
	resData := map[string]interface{}{
		"topics": serializers.GetAllFullTopicResponse(userModel.Topics),
	}

	c.JSON(http.StatusOK, serializers.SuccessBaseResponse(resData))
	return
}

func GetProfile(c *gin.Context) {
	userModel := c.MustGet("user_model").(models.User)
	if err := userModel.GetTopic(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	if err := userModel.GetTopicFollow(); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	c.JSON(http.StatusOK, serializers.SuccessBaseResponse(serializers.ProfileUserResponse(userModel)))
	return
}

func FollowTopic(c *gin.Context) {
	userModel := c.MustGet("user_model").(models.User)
	var v _interface.FollowTopic

	if err := c.Bind(&v); err != nil {
		err := err.(validator.ValidationErrors)

		errMes := models.CustomValidateError(err)
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity,errMes))
		return
	}

	if err := userModel.FollowTopic(v.TopicID); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ErrorBaseResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
	return
}
