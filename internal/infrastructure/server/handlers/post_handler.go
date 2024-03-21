package handlers

import (
	"net/http"

	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	uc *usecase.PostUseCase
}

func NewPostHandler(uc *usecase.PostUseCase) *PostHandler {
	return &PostHandler{uc}
}

func (h *PostHandler) CreatePost(ctx *gin.Context) {

	var req req.CreatePostReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

	form, _ := ctx.MultipartForm()
	user := ctx.GetStringMap("user")
	v, ok := user["userId"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   e.ErrKeyNotFound.Error(),
			Message: "Failed to get userId from token",
		})
		return
	}
	req.UserID = v.(string)
	req.Files = form.File["medias"]

	ctx.JSON(http.StatusCreated, h.uc.CreatePost(ctx, &req))
}
