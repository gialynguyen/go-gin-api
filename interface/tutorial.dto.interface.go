package _interface

type CreateTutorialDto struct {
	Name        string   `form:"name" json:"name" binding:"required,min=2,max=50"`
	Description string   `form:"desc" json:"desc" binding:"required,min=10,max=300"`
	TopicsId    []string `form:"topics_id" json:"topics_id" binding:"required,min=1"`
}

