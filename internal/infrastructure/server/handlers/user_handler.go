package handlers

import (
	"net/http"

	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var req req.SignUpReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": "failed to parse req",
			"error":  err.Error(),
		})
		return
	}
	user, err := h.uc.SignUp(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res.CommonRes{
			Code: http.StatusInternalServerError, Error: err.Error(), Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(200, res.CommonRes{
		Code:    200,
		Message: "Succesfully created user",
		Result:  user,
	})
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req req.LoginReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": "failed to parse req",
			"error":  err.Error(),
		})
		return
	}
	resp, err := h.uc.Login(ctx, &req)
	if err != nil {
		ctx.JSON(resp.Code, res.CommonRes{Code: resp.Code, Error: err.Error(), Message: resp.Message})
		return
	}
	ctx.JSON(resp.Code, resp)
}
