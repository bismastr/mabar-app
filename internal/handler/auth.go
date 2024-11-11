package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (h *Handler) Callback(ctx *gin.Context) {
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

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//Store Session
	session, _ := gothic.Store.Get(ctx.Request, "user_session")

	session.Values["user"] = user

	err = session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	sendSuccessResponse(ctx, user)
}

func (h *Handler) Login(ctx *gin.Context) {
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	if gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request); err == nil {
		sendSuccessResponse(ctx, gothUser)
		fmt.Println("err complete", err.Error())
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
