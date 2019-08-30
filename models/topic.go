package models

import (
	"fmt"
	"github.com/golang-gin/db"
	_interface "github.com/golang-gin/interface"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Topic struct {
	BaseModel
	Name           string `gorm:"column:name;not null"`
	Description    string `gorm:"column:desc;not null"`
	CreateByID     uuid.UUID
	CreateBy       *User
	FollowingUsers []TopicUser `gorm:"foreignkey:TopicId"`
	Tutorials      []Tutorial  `gorm:"many2many:topic_tutorial;"`
}

type TopicUser struct {
	BaseModel
	Topic   Topic
	TopicID uuid.UUID
	User    User
	UserID  uuid.UUID
}

func (*TopicUser) TableName() string {
	return "topic_user"
}

func CreateTopic(user User, data _interface.CreateTopicDto) (Topic, error) {
	_db := db.GetDB()

	var topic Topic

	if !_db.Where("name = ? AND create_by_id = ?", data.Name, user.ID).First(&topic).RecordNotFound() {
		return topic, errors.New("Topic has exists")
	}

	topic = Topic{Name: data.Name, Description: data.Description}

	topic.CreateByID = user.ID
	if err := topic.SetListTutorial(data.Tutorials); err != nil {
		return topic, err
	}

	if err := _db.Save(&topic).Error; err != nil {
		return topic, err
	}

	return topic, nil
}

func (t *Topic) SetListTutorial(tutorialsId []string) error {
	_db := db.GetDB()
	var tutorials []Tutorial
	for _, tutorialId := range tutorialsId {
		var tutorial Tutorial
		if _db.First(&tutorial, Tutorial{BaseModel: BaseModel{ID: uuid.FromStringOrNil(tutorialId)}}).RecordNotFound() {
			return errors.New("Tutorial has exists")
		}
		tutorials = append(tutorials, tutorial)

	}
	t.Tutorials = tutorials
	return nil
}

func (t *Topic) GetTopic() error {
	_db := db.GetDB()
	if _db.First(&t, Topic{BaseModel: BaseModel{ID: t.ID}}).RecordNotFound() {
		return errors.New("Topic not found")
	}
	var user User
	_db.Where("id = ?", t.CreateByID).First(&user)
	t.CreateBy = &user
	tx := _db.Begin()
	tx.Model(&t).Related(&t.Tutorials, "Tutorials")
	tx.Model(&t).Related(&t.FollowingUsers, "FollowingUsers")
	for i := range t.FollowingUsers {
		tx.Model(&t.FollowingUsers[i]).Related(&t.FollowingUsers[i].User)
	}
	for i := range t.Tutorials {
		tx.Model(&t.Tutorials[i]).Related(&t.Tutorials[i].Lessons, "Lessons")

	}
	err := tx.Commit().Error
	fmt.Println(t.CreateBy)
	return err
}
