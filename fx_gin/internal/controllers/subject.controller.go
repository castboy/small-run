package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"small-run/fx_gin/internal/models"
	"small-run/fx_gin/shared/middlewares"
)

type SubjectController struct {
	fx.Out
	Ctl IController `group:"router"`
}

func NewSubjectController(md *models.SubjectModel) SubjectController {
	return SubjectController{Ctl: &SubjectHandler{Md: md}}
}

type SubjectHandler struct {
	Md *models.SubjectModel
}

func (i *SubjectHandler) RegisterRoute(router gin.IRouter) {
	group := router.Group("/subject", i.Search, middlewares.NewJWT())
	group.POST("", i.Create)
	group.GET("/:id", i.Search)
}

func (i *SubjectHandler) Create(c *gin.Context) {
	err := i.Md.Create(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (i *SubjectHandler) Search(c *gin.Context) {
	res, err := i.Md.Search(c, "id")
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, res)
}

