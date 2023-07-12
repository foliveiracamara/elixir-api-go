package service

import (
	"github.com/foliveiracamara/elixir-api-go/src/config/app_error"
	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
	"go.uber.org/zap"
)

type PostDomainService interface {
	CreatePost(entity.PostDomainInterface) *app_error.AppErr
	UpdatePost(string, entity.PostDomainInterface) *app_error.AppErr
	FindPost(string) (*entity.PostDomainInterface, *app_error.AppErr)
	DeletePost(string) *app_error.AppErr
}

func NewPostDomainService() PostDomainService {
	return &postDomainService{}
}

type postDomainService struct{}

func (svc *postDomainService) CreatePost(postDomain entity.PostDomainInterface) *app_error.AppErr {
 
	logger.Info("Init createPost entity",
		zap.String("journey", "createPostService"),
	)
	postDomain.GenerateId()
	postDomain.GetPostDate()

	return nil
}

func (svc *postDomainService) DeletePost(string) *app_error.AppErr {
	return nil
}

func (svc *postDomainService) UpdatePost(string, entity.PostDomainInterface) *app_error.AppErr {
	return nil
}

func (svc *postDomainService) FindPost(string) (*entity.PostDomainInterface, *app_error.AppErr) {
	return nil, nil
}
