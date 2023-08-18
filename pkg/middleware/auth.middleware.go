package middleware

import (
	"net/http"
	"strings"

	"github.com/Bek0sh/market-place/pkg/config"
	"github.com/Bek0sh/market-place/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		token, err := ctx.Cookie("access_token")
		header := ctx.GetHeader("Autorization")

		fields := strings.Fields(header)

		if len(fields) != 0 || fields[0] == "Bearer" {
			accessToken = fields[0]
		} else if err != nil {
			accessToken = token
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  "fail",
					"message": "You are not logged in",
				},
			)
			return
		}

		config, _ := config.LoadConfig()

		_, err = utils.VerifyToken(accessToken, config.AccessTokenPublicKey)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  "fail",
					"message": "Access token is not valid",
				},
			)
			return
		}

		ctx.Next()

	}
}
