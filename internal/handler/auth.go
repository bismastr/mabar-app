package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Callback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to authenticate",
			"err":     err.Error(),
		})
		return
	}

	sendSuccessResponse(ctx, user)
}

func Login(ctx *gin.Context) {
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	if gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request); err == nil {
		sendSuccessResponse(ctx, gothUser)
	} else {
		fmt.Println("err complete userAuth", err.Error())
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	}
}

func sendSuccessResponse(ctx *gin.Context, user interface{}) {
	ctx.JSON(200, gin.H{
		"message": "Successfully authenticated",
		"user":    user,
	})
}
