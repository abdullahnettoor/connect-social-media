package handlers

import (
	"encoding/json"
	"net/http"

	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/kafka/producer"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	uc *usecase.ChatUseCase
}

func NewChatHandler(uc *usecase.ChatUseCase) *ChatHandler {
	return &ChatHandler{uc}
}

func (h *ChatHandler) SendChat(ctx *gin.Context) {

	var req req.SendChatReq
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "Failed to parse request",
		})
		return
	}

	user := ctx.GetStringMap("user")
	userId, ok := user["userId"]
	if !ok {
		ctx.JSON(http.StatusBadRequest, res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   e.ErrKeyNotFound.Error(),
			Message: "Failed to get userId from token",
		})
		return
	}
	req.SenderID = userId.(string)

	data, _ := json.Marshal(req)
	if err := producer.NewProducer("chat", req.RecipientID, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, res.CommonRes{
			Code:    http.StatusInternalServerError,
			Error:   err.Error(),
			Message: "Kafka Error",
		})
		return
	}

	resp := h.uc.SaveMessage(ctx, &req)
	ctx.JSON(resp.Code, resp)
}
