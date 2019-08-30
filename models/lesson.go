package models

import (
	"errors"
	"github.com/golang-gin/db"
	_interface "github.com/golang-gin/interface"
	uuid "github.com/satori/go.uuid"
)

type Lesson struct {
	BaseModel
	Name         string `gorm:"column:name;not null"`
	Description  string
	CreateByID   uuid.UUID
	CreateBy     *User
	TutorialID   uuid.UUID
	LearnedUsers []LessonUser `gorm:"foreignkey:LessonID;"`
}

type LessonUser struct {
	BaseModel
	Lesson   Lesson
	LessonID uuid.UUID
	User     User
	UserID   uuid.UUID
}

func CreateLesson(user *User, data _interface.CreateLessonDto) error {
	_db := db.GetDB()
	var tutorial Tutorial
	if _db.Where("id = ? AND create_by_id = ?", uuid.FromStringOrNil(data.TutorialId), user.ID).First(&tutorial).RecordNotFound() {
		return errors.New("Tutorial not found")
	}
	lesson := Lesson{
		Name:       data.Name,
		Description: data.Description,
		CreateByID: uuid.FromStringOrNil(data.TutorialId),
		CreateBy:   user,
		TutorialID: uuid.FromStringOrNil(data.TutorialId),
	}
	_db.Save(&lesson)

	return nil
}
