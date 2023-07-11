package service

import (
	"fmt"

	"github.com/foliveiracamara/elixir-api-go/src/config/app_error"
	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
	"github.com/foliveiracamara/elixir-api-go/src/domain"
	"go.uber.org/zap"
)

type DonorDomainService interface {
	ListDonors() 
	FindDonor(string) (*entity.DonorDomainInterface, *app_error.AppErr)
	CreateDonor(entity.DonorDomainInterface) *app_error.AppErr
	UpdateDonor(string, entity.DonorDomainInterface) *app_error.AppErr
	DeleteDonor(string) *app_error.AppErr
}

func NewDonorDomainService(repo domain.DonorRepositoryInterface) DonorDomainService {
	return &donorDomainService{
		repo: repo,
	}
}

type donorDomainService struct {
	repo domain.DonorRepositoryInterface
}

func (svc *donorDomainService) CreateDonor(donorDomain entity.DonorDomainInterface) *app_error.AppErr {
	logger.Info("Init createDonor entity", 
		zap.String("journey", "createDonorService"),
	)
	donorDomain.EncryptPassword()
	donorDomain.GenerateId()

	err := svc.repo.Create(donorDomain)
	if err != nil {
		panic(err)
	}

	return nil
}

func (svc *donorDomainService) DeleteDonor(string) *app_error.AppErr {
	return nil
}

func (svc *donorDomainService) UpdateDonor(string, entity.DonorDomainInterface) *app_error.AppErr {
	return nil
}

func (svc *donorDomainService) FindDonor(string) (*entity.DonorDomainInterface, *app_error.AppErr) {
	return nil, nil
}

func (svc *donorDomainService) ListDonors()  {
	for _, d := range Donors{
		fmt.Println(d)
	}
	// return Donors, nil
}

type donors *[]entity.DonorDomainInterface

var Donors []entity.DonorDomainInterface