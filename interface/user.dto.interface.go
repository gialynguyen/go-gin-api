package _interface

type CreateUserDto struct {
	Email     string `form:"email" json:"email" binding:"required,email"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type LoginUserDto struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type FollowTopic struct {
	TopicID string `form:"topic_id" json:"topic_id" binding:"required"`
}