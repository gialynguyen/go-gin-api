package _interface

type UserLoginResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type BaseInfoUserResponse struct {
	ID string `json:"id"`
	UserRegisterResponse
}

type ProfileUserResponse struct {
	BaseInfoUserResponse
	TotalTopicFollow int             `json:"total_topic_follow"`
	TopicFollow      []TopicResponse `json:"topic_follow"`
	Topics           []TopicResponse `json:"topics"`
}
