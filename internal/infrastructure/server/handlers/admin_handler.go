package handlers

import (
	"net/http"

	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	uc *usecase.AdminUseCase
}

func NewAdminHandler(uc *usecase.AdminUseCase) *AdminHandler {
	return &AdminHandler{uc: uc}
}

func (h *AdminHandler) Login(ctx *gin.Context) {
	var req req.AdminLoginReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

	resp := h.uc.Login(ctx, &req)
	if resp.Error != nil {
		ctx.JSON(resp.Code, res.CommonRes{
			Code:    resp.Code,
			Error:   resp.Error,
			Message: resp.Message,
		})
		return
	}
	ctx.JSON(resp.Code, resp)
}

func (h *AdminHandler) BlockUser(ctx *gin.Context) {
	var req req.UserId
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

	resp := h.uc.BlockUser(ctx, &req)

	ctx.JSON(resp.Code, resp)
}

func (h *AdminHandler) UnblockUser(ctx *gin.Context) {
	var req req.UserId
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": "failed to parse req",
			"error":  err.Error(),
		})
		return
	}

	resp := h.uc.UnblockUser(ctx, &req)

	ctx.JSON(resp.Code, resp)
}
