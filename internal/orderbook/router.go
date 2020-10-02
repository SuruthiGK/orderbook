package orderbook

import (
	"orderbook/internal/orderbook/api"

	"github.com/gin-gonic/gin"
)

func RouterMain(router *gin.Engine) {
	api.Routes(router)
}
