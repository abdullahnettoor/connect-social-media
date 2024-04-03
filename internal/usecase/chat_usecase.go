package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	"github.com/google/uuid"
)

type ChatUseCase struct {
	chatRepo *repo.ChatRepository
}

func NewChatUseCase(chatRepo *repo.ChatRepository) *ChatUseCase {
	return &ChatUseCase{chatRepo}
}

func (uc *ChatUseCase) SaveMessage(ctx context.Context, req *req.SendChatReq) *res.CommonRes {
	msg := &entity.Message{
		ID:          uuid.NewString(),
		SenderID:    req.SenderID,
		RecipientID: req.RecipientID,
		Message:     req.Message,
		CreatedAt:   helper.CurrentIsoDateTimeString(),
	}

	fmt.Println("Current time is", helper.CurrentIsoDateTimeString())

	msg, err := uc.chatRepo.CreateMessage(ctx, msg)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "Error saving message to db",
			Error:   err.Error(),
		}
	}

	return &res.CommonRes{
		Code:    http.StatusOK,
		Message: "Message sent successful",
		Result:  msg,
	}
}
