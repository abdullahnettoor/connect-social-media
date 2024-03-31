package usecase

import (
	"context"
	"net/http"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	"github.com/google/uuid"
)

type CommentUseCase struct {
	commentRepo *repo.CommentRepository
	contentRepo *repo.ContentRepository
}

func NewCommentUseCase(commentRepo *repo.CommentRepository, contentRepo *repo.ContentRepository) *CommentUseCase {
	return &CommentUseCase{commentRepo, contentRepo}
}

func (uc *CommentUseCase) CreateComment(ctx context.Context, req *req.CreateCommentReq) *res.CommonRes {

	var comment = &entity.Comment{
		ID:        uuid.NewString(),
		Comment:   req.Comment,
		CreatedAt: helper.CurrentIsoDateTimeString(),
		PostID:    req.PostID,
		UserID:    req.UserID,
	}

	comment, err := uc.commentRepo.Create(ctx, comment)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "Db error",
			Error:   err.Error(),
		}
	}

	return &res.CommonRes{
		Code:    http.StatusCreated,
		Message: "Added New Comment",
		Result:  comment,
	}
}

func (uc *CommentUseCase) DeleteComment(ctx context.Context, req *req.DeleteCommentReq) *res.CommonRes {

	err := uc.commentRepo.Delete(ctx, req.UserID, req.CommentID)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "Db error",
			Error:   err.Error(),
		}
	}

	return &res.CommonRes{
		Code:    http.StatusOK,
		Message: "Delete Comment Successful",
	}
}

func (uc *CommentUseCase) GetCommentsByPostId(ctx context.Context, req *req.GetCommentsReq) *res.GetCommentsRes {
	comments, err := uc.commentRepo.GetCommentsOfPost(ctx, req.PostID)
	if err != nil {
		return &res.GetCommentsRes{CommonRes: res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "Error retrieving posts",
			Error:   err.Error(),
		}}
	}

	return &res.GetCommentsRes{CommonRes: res.CommonRes{
		Code:    http.StatusOK,
		Message: "Fetched All posts"},
		Comments: comments,
	}
}
