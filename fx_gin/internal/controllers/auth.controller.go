package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"small-run/fx_gin/internal/models"
	"small-run/fx_gin/shared/tools"
)

type AuthController struct {
	fx.Out
	Ctl IController `group:"router"`
}

func NewAuthController(md *models.AuthModel) AuthController {
	return AuthController{Ctl: &AuthHandler{Md: md}}
}

type AuthHandler struct {
	Md *models.AuthModel
}

func (i *AuthHandler) RegisterRoute(router gin.IRouter) {
	group := router.Group("/auth")
	group.GET("", i.CheckAuth)
}

func (i *AuthHandler) CheckAuth(c *gin.Context) {
	ok, err := i.Md.CheckAuth(c)
	fmt.Println(ok)
	if err != nil {
		return
	}

	if ok {
		username := c.Query("username")
		password := c.Query("password")

		token, err := tools.GenerateToken(username, password)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}

}