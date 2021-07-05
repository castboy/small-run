package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"time"
)

type DealTime struct {
	fx.Out
	Hd gin.HandlerFunc `group:"handleFunc"`
}

func NewDealTime(log *log.Logger) DealTime {
	hd := func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 统计时间
		since := time.Since(start)
		log.Println("程序用时：", since)
	}

	return DealTime{Hd: hd}
}
