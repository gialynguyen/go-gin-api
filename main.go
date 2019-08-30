package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-gin/config"
	"github.com/golang-gin/controller"
	"github.com/golang-gin/db"
	"github.com/golang-gin/models"

	"github.com/jinzhu/gorm"
	_ "net/http"
	"os"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Tutorial{}, &models.Topic{}, &models.TopicUser{}, &models.TutorialUser{}, &models.Lesson{}, &models.LessonUser{})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("api/v1")
	controller.UserRouter(api)
	controller.TopicRouter(api)
	controller.TutorialRouter(api)
	controller.LessonRouter(api)
	return router
}

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	db.Init()

	Migrate(db.DB)

	r := SetupRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
