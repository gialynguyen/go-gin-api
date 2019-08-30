package _interface

type TutorialResponse struct {
	ID          string           `json:"id"`
	Description string           `json:"desc"`
	Name        string           `json:"name"`
	TotalLesson int              `json:"total_lesson"`
	Lessons     []LessonResponse `json:"lessons"`
}
