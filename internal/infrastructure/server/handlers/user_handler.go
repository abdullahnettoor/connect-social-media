package handlers

import (
	"fmt"
	"net/http"

	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	var req req.SignUpReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

	resp := h.uc.SignUp(ctx, &req)
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

func (h *UserHandler) Login(ctx *gin.Context) {
	var req req.LoginReq
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

func (h *UserHandler) VerifyOtp(ctx *gin.Context) {
	var req req.VerifyOtp
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

	user := ctx.GetStringMap("user")
	fmt.Println("User is", user)
	v, ok := user["userId"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   e.ErrKeyNotFound.Error(),
			Message: "Failed to parse request",
		})
		return
	}
	req.UserID = v.(string)

	resp := h.uc.VerifyOtp(ctx, &req)
	ctx.JSON(resp.Code, resp)
}
