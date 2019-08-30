package models

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang-gin/config"
	"github.com/golang-gin/db"
	"github.com/golang-gin/interface"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	BaseModel
	Username         string         `gorm:"column:username"`
	Email            string         `gorm:"column:email;unique"`
	PasswordHash     string         `gorm:"column:password;not null"`
	Topics           []Topic        `gorm:"foreignkey:CreateByID"`
	Tutorials        []Tutorial     `gorm:"foreignkey:CreateByID"`
	Lessons          []Lesson       `gorm:"foreignkey:CreateByID"`
	TopicsFollowing  []TopicUser    `gorm:"foreignkey:UserID;"`
	TutorialLearning []TutorialUser `gorm:"foreignkey:UserID;"`
	LearnedLesson    []LessonUser   `gorm:"foreignkey:UserID;"`
}

type JwtToken struct {
	Username string    `json:"user_name"`
	Id       uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

const (
	tokenExp = 30
)

func (u *User) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not empty")
	}
	bytePassword := []byte(password)
	hashPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(hashPassword)
	return nil
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindOne(condition _interface.LoginUserDto) (User, error) {
	_db := db.GetDB()
	user := User{Email: condition.Email}
	err := _db.Where(user).First(&user).Error
	return user, err
}

func CreateUser(data _interface.CreateUserDto) (User, error) {
	_db := db.GetDB()
	user := User{
		Email:    data.Email,
		Username: fmt.Sprintf(data.FirstName + data.LastName),
	}
	_ = user.setPassword(data.Password)

	if err := _db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GeneratorAccessToken(user User) string {
	_secretKey := config.GetConfig().GetString("jwtaccesstokenkey")

	exp := time.Now().Add(tokenExp * time.Minute)

	claims := &JwtToken{
		user.Username,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := jwtToken.SignedString([]byte(_secretKey))

	return accessToken
}

func (u *User) GetTopic() error {
	_db := db.GetDB()
	tx := _db.Begin()
	tx.Model(&u).Related(&u.Topics, "Topics")
	for i := range u.Topics {
		u.Topics[i].CreateBy = u
		_ = u.Topics[i].GetTopic()
		tx.Model(&u.Topics[i]).Related(&u.Topics[i].Tutorials, "Tutorials")
		for v := range u.Topics[i].Tutorials {
			tx.Model(&u.Topics[i].Tutorials[v]).Related(&u.Topics[i].Tutorials[v].Lessons)
		}
	}
	err := tx.Commit().Error
	return err
}


func (u *User) GetTutorial() {
	_db := db.GetDB()
	_db.Model(&u).Related(&u.Tutorials, "Tutorials")

}

func (u *User) FollowTopic(topicID string) error{
	_db := db.GetDB()
	topic := Topic{BaseModel: BaseModel{ID: uuid.FromStringOrNil(topicID)}}
	if _db.Where(&topic).First(&topic).RecordNotFound() {
		return errors.New("Topic not found ")
	}
	if topic.CreateByID == u.ID {
		return errors.New("You can't follow this topic ")
	}

	u.TopicsFollowing = append(u.TopicsFollowing, TopicUser{
		Topic: topic,
		TopicID: topic.ID,
		User: *u,
		UserID: u.ID,
	})

	_db.Save(&u)

	return nil
}

func (u *User) GetTopicFollow() error {
	_db := db.GetDB()
	tx := _db.Begin()
	tx.Model(&u).Related(&u.TopicsFollowing, "TopicsFollowing")
	for i := range u.TopicsFollowing {
		tx.Model(&u.TopicsFollowing[i]).Related(&u.TopicsFollowing[i].Topic)
	}
	err := tx.Commit().Error

	return err
}