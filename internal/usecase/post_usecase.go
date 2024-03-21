package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	"github.com/google/uuid"
)

type PostUseCase struct {
	postRepo    *repo.PostRepository
	contentRepo *repo.ContentRepository
}

func NewPostUseCase(postRepo *repo.PostRepository, contentRepo *repo.ContentRepository) *PostUseCase {
	return &PostUseCase{postRepo, contentRepo}
}

func (uc *PostUseCase) CreatePost(ctx context.Context, req *req.CreatePostReq) *res.CommonRes {

	var wg sync.WaitGroup
	urlChannel := make(chan string, len(req.Files))
	fmt.Println("", req.Files)
	for i, file := range req.Files {
		fmt.Println("File no:", i)
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			dir := "posts/"
			fileType := file.Header.Get("Content-Type")
			switch fileType {
			case "image/jpeg", "image/png", "image/webp":
				dir += "images"
			case "video/mp4", "video/avi", "video/webm", "video/quicktime":
				dir += "videos"
			}

			contents, err := file.Open()
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer contents.Close()

			url, err := uc.contentRepo.UploadImage(ctx, uuid.NewString(), dir, contents)
			if err != nil {
				fmt.Println("Error uploading file:", err)
				return
			}

			urlChannel <- url
		}(file)
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	close(urlChannel)
	fmt.Println("All goroutines have finished.")

	urls := make([]string, 0, len(req.Files))
	for url := range urlChannel {
		urls = append(urls, url)
	}
	fmt.Println("😀")

	var post = &entity.Post{
		ID:          uuid.NewString(),
		Description: req.Description,
		Location:    req.Location,
		MediaUrls:   urls,
		IsBlocked:   false,
		CreatedAt:   helper.CurrentIsoDateTimeString(),
		UpdatedAt:   helper.CurrentIsoDateTimeString(),
	}

	post, err := uc.postRepo.Create(ctx, req.UserID, post)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "Db error",
			Error:   err.Error(),
		}
	}

	return &res.CommonRes{
		Code:    http.StatusCreated,
		Message: "Created New Post",
		Result:  post,
	}
}
