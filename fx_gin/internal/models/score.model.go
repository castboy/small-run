package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"small-run/fx_gin/internal/models/entities"
	"small-run/fx_gin/shared/tools"
)

type ScoreModel struct {
	db *gorm.DB
}

func NewScoreModel(db *gorm.DB) (model *ScoreModel, err error) {
	return &ScoreModel{db: db}, db.AutoMigrate(&entities.ScoreEntity{})
}

func (i *ScoreModel) Create (c *gin.Context) (*entities.ScoreEntity, error) {
	var body entities.ScoreEntity
	if err := c.ShouldBindJSON(&body); err != nil {
		return nil, c.Error(err)
	}

	if err := i.db.Omit("ID").Create(&body).Error; err != nil {
		return nil, c.Error(err)
	}

	return &body, nil
}

func (i *ScoreModel) Search (c *gin.Context) (*entities.ScoreEntity, error) {
	var (
		student_id, subject_id  int
		err  error
		body entities.ScoreEntity
	)
	if student_id, err = tools.GetPathInt(c, "student_id"); err != nil {
		return nil, c.Error(err)
	}

	if subject_id, err = tools.GetPathInt(c, "subject_id"); err != nil {
		return nil, c.Error(err)
	}

	if err = i.db.Where("student_id=? and subject_id=?", student_id, subject_id).First(&body).Error; err != nil {
		return nil, c.Error(err)
	}

	return &body, nil
}


