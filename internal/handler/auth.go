package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bismastr/discord-bot/internal/auth"
	"github.com/bismastr/discord-bot/internal/config"
	"github.com/bismastr/discord-bot/internal/repository"
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
			"message": "Failed to authenticate complete user auth",
			"err":     err.Error(),
		})
		return
	}

	discordUid, _ := strconv.ParseInt(user.UserID, 10, 64)
	err = h.user.Createuser(ctx, repository.InsertUserParams{
		Username:   user.Name,
		AvatarUrl:  user.AvatarURL,
		DiscordUid: discordUid,
	})

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to create user to db",
			"err":     err.Error(),
		})
	}

	err = h.auth.StoreUserSession(ctx.Writer, ctx.Request, user)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, config.Envs.CallbackRedirectUrl)
}

func (h *Handler) Login(ctx *gin.Context) {
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	if gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request); err == nil {
		sendSuccessResponse(ctx, gothUser)

	} else {
		fmt.Println("not found sesson")
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	}
}

func (h *Handler) CheckIsAuthenticaed(ctx *gin.Context) {
	u, err := h.auth.GetUserSession(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusOK, auth.User{})
		return
	}

	user := &auth.User{
		Name:      u.Name,
		UserID:    u.UserID,
		AvatarURL: u.AvatarURL,
	}

	ctx.JSON(http.StatusOK, user)
}

func sendSuccessResponse(ctx *gin.Context, user interface{}) {
	ctx.JSON(200, gin.H{
		"message": "Successfully authenticated",
		"user":    user,
	})
}
