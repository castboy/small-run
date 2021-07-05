package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"small-run/fx_gin/internal/models"
	"small-run/fx_gin/shared/middlewares"
)

type StudentController struct {
	fx.Out
	Ctl IController `group:"router"`
}

func NewStudentController(md *models.StudentModel) StudentController {
	return StudentController{Ctl: &StudentHandler{Md: md}}
}

type StudentHandler struct {
	Md *models.StudentModel
}

func (i *StudentHandler) RegisterRoute(router gin.IRouter) {
	group := router.Group("/student", i.Search, middlewares.NewJWT())
	group.POST("", i.Create)
	group.GET("/:id", i.Search)
}

func (i *StudentHandler) Create(c *gin.Context) {
	body, err := i.Md.Create(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, body)
}

func (i *StudentHandler) Search(c *gin.Context) {
	res, err := i.Md.Search(c, "id")
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, res)
}

