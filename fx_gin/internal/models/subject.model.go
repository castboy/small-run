package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"small-run/fx_gin/internal/models/entities"
	"small-run/fx_gin/shared/tools"
)

type SubjectModel struct {
	db *gorm.DB
}

func NewSubjectModel(db *gorm.DB) (model *SubjectModel, err error) {
	return &SubjectModel{db: db}, db.AutoMigrate(&entities.SubjectEntity{})
}

func (i *SubjectModel) Create (c *gin.Context) error {
	var body entities.SubjectEntity
	if err := c.ShouldBindJSON(&body); err != nil {
		return c.Error(err)
	}

	if err := i.db.Omit("ID").Create(&body).Error; err != nil {
		return c.Error(err)
	}

	return nil
}



func (i *SubjectModel) Search (c *gin.Context, field string) (*entities.SubjectEntity, error) {
	var (
		id int
		err error
		body entities.SubjectEntity
	)

	if id, err = tools.GetPathInt(c, field); err != nil {
		return nil, c.Error(err)
	}

	if err := i.db.First(&body, id).Error; err != nil {
		return nil, c.Error(err)
	}

	return &body, nil
}


