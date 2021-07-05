package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"small-run/fx_gin/internal/models/entities"
	"small-run/fx_gin/shared/tools"
)

type AuthModel struct {
	db *gorm.DB
}

func NewAuthModel(db *gorm.DB) (model *AuthModel, err error) {
	return &AuthModel{db: db}, db.AutoMigrate(&entities.AuthEntity{})
}

func (i *AuthModel) CheckAuth (c *gin.Context) (bool, error) {
	var body entities.AuthEntity

	username := c.Query("username")
	password := c.Query("password")

	err := i.db.Select("id").Where(entities.AuthEntity{Username: username, Password: tools.EncodeMD5(password)}).First(&body).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if body.ID > 0 {
		return true, nil
	}

	return false, nil
}
