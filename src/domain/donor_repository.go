package domain

import "github.com/foliveiracamara/elixir-api-go/src/domain/entity"

type DonorRepositoryInterface interface {
	Create(donor entity.DonorDomainInterface) (error)
}