package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

const (
	SessionName = "_user_session"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := gothic.Store.Get(ctx.Request, SessionName)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
		}

		u := session.Values["user"]

		ctx.Set("user", u)
		ctx.Next()
	}

}
