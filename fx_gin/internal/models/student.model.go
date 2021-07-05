package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"small-run/fx_gin/internal/models/entities"
	"small-run/fx_gin/shared/tools"
)

type StudentModel struct {
	db *gorm.DB
}

func NewStudentModel(db *gorm.DB) (model *StudentModel, err error) {
	return &StudentModel{db: db}, db.AutoMigrate(&entities.StudentEntity{})
}

func (i *StudentModel) Create (c *gin.Context) (*entities.StudentEntity, error) {
	var body entities.StudentEntity
	if err := c.ShouldBindJSON(&body); err != nil {
		return nil, c.Error(err)
	}

	if err := i.db.Omit("ID").Create(&body).Error; err != nil {
		return nil, c.Error(err)
	}

	return &body, nil
}

func (i *StudentModel) Search (c *gin.Context, field string) (*entities.StudentEntity, error) {
	var (
		id int
		err error
		body entities.StudentEntity
	)

	if id, err = tools.GetPathInt(c, field); err != nil {
		return nil, c.Error(err)
	}

	if err := i.db.First(&body, id).Error; err != nil {
		return nil, c.Error(err)
	}

	return &body, nil
}


