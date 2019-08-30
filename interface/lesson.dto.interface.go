package _interface

type CreateLessonDto struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=100"`
	Description string `form:"desc" json:"desc" binding:"required,min=10,max=300"`
	TutorialId string `form:"tutorial_id" json:"tutorial_id" binding:"required"`
}
