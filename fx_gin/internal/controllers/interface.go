package controllers

import "github.com/gin-gonic/gin"

type IController interface {
	RegisterRoute(router gin.IRouter)
}
