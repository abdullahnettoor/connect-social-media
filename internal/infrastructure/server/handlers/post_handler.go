package handlers

import (
	"net/http"

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
	req.UserID = int64(user["userId"].(float64))
	req.Files = form.File["images"]

	ctx.JSON(http.StatusCreated, h.uc.CreatePost(ctx, &req))
}
