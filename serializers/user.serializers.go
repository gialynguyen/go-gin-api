package serializers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-gin/interface"
	"github.com/golang-gin/models"
)

func LoginResponse(C *gin.Context) _interface.UserLoginResponse {
	userModel := C.MustGet("user_model").(models.User)

	return _interface.UserLoginResponse{Email: userModel.Email, Username: userModel.Username, AccessToken: models.GeneratorAccessToken(userModel)}
}

func RegisterResponse(user models.User) _interface.UserRegisterResponse {
	resData := _interface.UserRegisterResponse{Email: user.Email, Username: user.Username}

	return resData
}

func BaseInforUserResponse(user models.User) _interface.BaseInfoUserResponse {
	return _interface.BaseInfoUserResponse{ID: user.ID.String(), UserRegisterResponse: _interface.UserRegisterResponse{Email: user.Email, Username: user.Username}}
}

func ProfileUserResponse(user models.User) _interface.ProfileUserResponse {
	followTopic := make([]models.Topic, len(user.TopicsFollowing))
	for i, t := range user.TopicsFollowing {
		followTopic[i] = t.Topic
	}

	return _interface.ProfileUserResponse{
		BaseInfoUserResponse: BaseInforUserResponse(user) ,
		Topics: GetAllFullTopicResponse(user.Topics),
		TotalTopicFollow: len(user.TopicsFollowing),
		TopicFollow: GetAllFullTopicResponse(followTopic),
	}
}