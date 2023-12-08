package middlewares

import (
	"strings"

	"github.com/bhoopendrau/tailscale-ui-backend/config"
	"github.com/gin-gonic/gin"
)

type GeneralResponse struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqKey := c.Request.Header.Get("X-Auth-Key")
		reqSecret := c.Request.Header.Get("X-Auth-Secret")
		configKey := config.GetString("http.auth.key")
		configSecret := config.GetString("http.auth.secret")

		if len(strings.TrimSpace(reqKey)) == 0 {
			c.JSON(401, GeneralResponse{Message: "Please provide Auth key", Reason: "Unauthorised"})
			return
		}
		if len(strings.TrimSpace(reqSecret)) == 0 {
			c.JSON(401, GeneralResponse{Message: "Please provide Auth secret", Reason: "Unauthorised"})
			return
		}
		if configKey != reqKey || configSecret != reqSecret {
			c.JSON(401, GeneralResponse{Message: "Invalid credentials", Reason: "Unauthorised"})
			return
		}
		c.Next()
	}
}
