package port

import (
	"database/sql"

	"github.com/foliveiracamara/elixir-api-go/src/domain"
	"github.com/foliveiracamara/elixir-api-go/src/port/repository/entity"
)

// Repositories contains all the repo structs
type Repositories struct {
	DonorRepository domain.DonorRepositoryInterface
}

// InitDonorRepository initializes the donors repository 
func InitDonorRepository(db *sql.DB) *Repositories {
	donorRepository := donor_repository.NewDonorRepository(db)
	return &Repositories{
		DonorRepository: donorRepository,
	}
}
