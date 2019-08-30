package serializers

import (
	_interface "github.com/golang-gin/interface"
	"github.com/golang-gin/models"
)

func LessonOfTutorial(tutorial models.Tutorial) []_interface.LessonResponse {
	jsonClasses := make([]_interface.LessonResponse, len(tutorial.Lessons))
	for i, v := range tutorial.Lessons {
		lesson := _interface.LessonResponse{
			ID: v.ID.String(),
			Name: v.Name,
			Description: v.Description,
		}
		jsonClasses[i] = lesson
	}
	return jsonClasses
}