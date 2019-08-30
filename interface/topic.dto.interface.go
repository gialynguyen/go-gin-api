package _interface

type CreateTopicDto struct {
	Name string  `form:"name" json:"name" binding:"required,min=2,max=50"`
	Description string `form:"desc" json:"desc" binding:"required,min=10,max=300"`
	Tutorials []string `form:"tutorials" json:"tutorials"`
}

type GetAllTopicDto struct {
	TopicID string `form:"topic_id" json:"topic_id" binding:"required"`
}
