package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v8"
)

type BaseModel struct {
	ID uuid.UUID `gorm:"type: uuid;primary_key"`
}

func (u *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

func CustomValidateError(err validator.ValidationErrors) map[string]interface{} {
	errMes := make(map[string]interface{})
	for _, v := range err {
		errMes[v.Field] = fmt.Sprintf("%v field must: %v", v.Field, v.Tag)
	}
	return errMes
}
