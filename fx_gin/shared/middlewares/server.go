package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MiddleWares struct {
	fx.In
	Hds []gin.HandlerFunc `group:"handleFunc"`
}

func Register(eng *gin.Engine, mds MiddleWares) {
	for i := range mds.Hds {
		eng.Use(mds.Hds[i])
	}
}
