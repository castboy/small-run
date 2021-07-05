package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Controllers struct {
	fx.In
	ICs []IController `group:"router"`
}

func Register(router gin.IRouter, c Controllers) {
	for i := range c.ICs {
		c.ICs[i].RegisterRoute(router)
	}
}
