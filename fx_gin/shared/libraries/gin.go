package libraries

import "github.com/gin-gonic/gin"

func NewEngine() *gin.Engine {
	return gin.Default()
}
