package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-gin/config"
	"github.com/golang-gin/db"
	"github.com/golang-gin/models"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func UpdateContextUserModel(c *gin.Context, userId uuid.UUID) error {

	user := models.User{BaseModel: models.BaseModel{ID: userId}}
	_db := db.GetDB()
	err := _db.First(&user, user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return err
	}

	c.Set("user_id", userId)
	c.Set("user_model", user)
	return nil
}

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequestWithClaims(c.Request, request.OAuth2Extractor, &models.JwtToken{}, func(token *jwt.Token) (interface{}, error) {
			_secretKey := config.GetConfig().GetString("jwtaccesstokenkey")
			b := []byte(_secretKey)
			return b, nil
		})

		if err != nil {
			if errDetail, ok := err.(*jwt.ValidationError); ok {
				if errDetail.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					c.JSON(http.StatusUnauthorized, gin.H{
						"error": "Token has Expired",
					})
					c.Abort()
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
				c.Abort()
			}
			c.Abort()
		}

		if claims, ok := token.Claims.(*models.JwtToken); ok && token.Valid {
			userId := claims.Id

			if err = UpdateContextUserModel(c, userId); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
				c.Abort()

			}

		}
		c.Next()
	}

}
