package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"small-run/fx_gin/internal/dto"
	"small-run/fx_gin/internal/models"
	"small-run/fx_gin/shared/middlewares"
)

type ScoreController struct {
	fx.Out
	Ctl IController `group:"router"`
}

func NewScoreController(sub *models.SubjectModel, stu *models.StudentModel, score *models.ScoreModel) ScoreController {
	return ScoreController{Ctl: &ScoreHandler{Sub: sub, Stu: stu, Score: score}}
}

type ScoreHandler struct {
	Sub *models.SubjectModel
	Stu *models.StudentModel
	Score *models.ScoreModel
}

func (i *ScoreHandler) RegisterRoute(router gin.IRouter) {
	group := router.Group("/score", middlewares.NewJWT())
	group.POST("", i.Create)
	group.GET("/:subject_id/:student_id", i.Search)
}

func (i *ScoreHandler) Create(c *gin.Context) {
	body, err := i.Score.Create(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, body)
}

func (i *ScoreHandler) Search(c *gin.Context) {
	sub, err := i.Sub.Search(c, "subject_id")
	if err != nil {
		return
	}

	stu, err := i.Stu.Search(c, "student_id")
	if err != nil {
		return
	}

	score, err := i.Score.Search(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, dto.BeRes(stu, sub, score))
}

