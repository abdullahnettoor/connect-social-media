package repo

import (
	"context"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByUserId(ctx context.Context, id string) (*entity.User, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.User, error)
	FollowUser(ctx context.Context, userId string, followedId string) error
	GetFollowers(ctx context.Context, userId string) ([]*entity.User, error)
	GetFollowing(ctx context.Context, userId string) ([]*entity.User, error)
	RemoveUserByEmail(ctx context.Context, email string) error
	UnfollowUser(ctx context.Context, userId string, followedId string) error
	UpdateUserStatus(ctx context.Context, id string, status string, updatedAt string) (*entity.User, error)
}

type ChatRepositoryInterface interface {
	CreateMessage(ctx context.Context, message *entity.Message) (*entity.Message, error)
}
