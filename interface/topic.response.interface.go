package _interface

type TopicResponse struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	CreateBy           BaseInfoUserResponse   `json:"create_by"`
	Description        string                 `json:"desc"`
	TotalTutorial      int                    `json:"total_tutorial"`
	Tutorials          []TutorialResponse     `json:"tutorials"`
	TotalUserFollowing int                    `json:"total_user_following"`
	UserFollowing      []BaseInfoUserResponse `json:"user_following"`
}
