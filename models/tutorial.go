package models

import (
	"github.com/golang-gin/db"
	_interface "github.com/golang-gin/interface"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Tutorial struct {
	BaseModel
	Name          string  `gorm:"column:name;not null"`
	Description   string  `gorm:"column:desc;not null"`
	Topics        []Topic `gorm:"many2many:topic_tutorial;"`
	CreateByID    uuid.UUID
	CreateBy      User
	Lessons       []Lesson       `gorm:"foreignkey:TutorialID;"`
	LearningUsers []TutorialUser `gorm:"foreignkey:TutorialID;"`
}

type TutorialUser struct {
	BaseModel
	Tutorial   Tutorial
	TutorialID uuid.UUID
	User       User
	UserID     uuid.UUID
	Owner      bool
	HasEnd     bool
	HasLearn   int
}

func CreateTutorial(user *User, data _interface.CreateTutorialDto) error {
	_db := db.GetDB()
	var tutorial Tutorial

	if !_db.Where("name = ? AND create_by_id = ?", data.Name, user.ID).First(&tutorial).RecordNotFound() {
		return errors.New("Tutorial has exists")
	}

	tutorial = Tutorial{Name: data.Name, Description: data.Description, CreateByID: user.ID}
	var topics []Topic

	for _, topicId := range data.TopicsId {
		var topic Topic
		if _db.Where("id = ? AND create_by_id = ?", topicId, user.ID).First(&topic).RecordNotFound() {
			return errors.New("Topic not found")
		} else {
			topics = append(topics, topic)
		}
	}
	tutorial.Topics = topics
	_db.Save(&tutorial)
	return nil
}

func (t *Tutorial) GetAllLesson() error {
	_db := db.GetDB()

	if _db.Where("id = ?", t.ID).First(&t).RecordNotFound() {
		return errors.New("Tutorial not found")
	}

	_db.Model(&t).Related(&t.Lessons, "Lessons")
	return nil
}
