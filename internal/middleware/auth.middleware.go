package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/EdwardKerckhof/gohtmx/pkg/response"
	"github.com/EdwardKerckhof/gohtmx/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	headerKey  = "authorization"
	typeBearer = "bearer"
	PayloadKey = "authPayload"
)

// TODO: unit tests (backend #22)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerKey := ctx.GetHeader(headerKey)
		if len(headerKey) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(errors.New("no authorization header provided")))
			return
		}

		fields := strings.Fields(headerKey)
		if len(fields) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(errors.New("invalid authorization header")))
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != typeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(errors.New("unsupported authorization type")))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(err))
			return
		}

		ctx.Set(PayloadKey, payload)
		ctx.Next()
	}
}
