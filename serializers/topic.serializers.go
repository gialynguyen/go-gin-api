package serializers

import (
	_interface "github.com/golang-gin/interface"
	"github.com/golang-gin/models"
)

func GetFullTopicResponse(topic models.Topic) _interface.TopicResponse {
	tutorials := GetAllTutorialOfTopic(topic.Tutorials)
	userFollow := make([]_interface.BaseInfoUserResponse,len(topic.FollowingUsers))

	for i, u := range topic.FollowingUsers {
		userFollow[i] = BaseInforUserResponse(u.User)
	}

	resData := _interface.TopicResponse{
		ID:                 topic.ID.String(),
		CreateBy: 			BaseInforUserResponse(*topic.CreateBy),
		Description:        topic.Description,
		Name:               topic.Name,
		Tutorials:          tutorials,
		UserFollowing:		userFollow,
		TotalTutorial:      len(topic.Tutorials),
		TotalUserFollowing: len(topic.FollowingUsers),
	}
	return resData
}

func GetAllFullTopicResponse(topics []models.Topic) []_interface.TopicResponse {
	resData := make([]_interface.TopicResponse, len(topics))

	for i, t := range topics {
		topicRes := GetFullTopicResponse(t)
		resData[i] = topicRes
	}

	return resData
}

func GetAllTutorialOfTopic(tutorials []models.Tutorial) []_interface.TutorialResponse {
	resData := make([]_interface.TutorialResponse, len(tutorials))

	for i, t := range tutorials {
		lessons := LessonOfTutorial(t)
		tutorial := _interface.TutorialResponse{
			ID:          t.ID.String(),
			Name:        t.Name,
			Description: t.Description,
			Lessons:     lessons,
			TotalLesson: len(lessons),
		}
		resData[i] = tutorial
	}
	return resData
}
