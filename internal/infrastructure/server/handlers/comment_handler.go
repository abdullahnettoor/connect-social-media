package handlers

import (
	"net/http"

	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	uc *usecase.CommentUseCase
}

func NewCommentHandler(uc *usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{uc}
}

func (h *CommentHandler) CreateComment(ctx *gin.Context) {

	var req req.CreateCommentReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

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

	resp := h.uc.CreateComment(ctx, &req)
	if resp.Error != nil {
		ctx.JSON(resp.Code, resp)
		return
	}

	ctx.JSON(resp.Code, resp)
}

func (h *CommentHandler) DeleteComment(ctx *gin.Context) {

	var req req.DeleteCommentReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

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

	resp := h.uc.DeleteComment(ctx, &req)

	ctx.JSON(resp.Code, resp)
}
