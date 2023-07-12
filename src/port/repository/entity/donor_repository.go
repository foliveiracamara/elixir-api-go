package donor_repository

import (
	"database/sql"
	"fmt"

	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
)

type DonorRepository struct {
	db *sql.DB
}

func NewDonorRepository(db *sql.DB) *DonorRepository {
	return &DonorRepository{
		db: db,
	}
}

// CreateDonorRepository creates a new donor in the db..
func (r *DonorRepository) Create(donor entity.DonorDomainInterface) (error) {

	stmt := `INSERT INTO tab_donor(
		donor_id
		, name
		, cpf
		, email
		, password
		, gender
		, blood_type
		, organ_donor
		, donation_date
		, birth_date) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		
	result, err := r.db.Exec(
		stmt, 
		donor.GetDonorId(),
		donor.GetName(),
		donor.GetCpf(),
		donor.GetEmail(),
		donor.GetPassword(),
		donor.GetGender(),
		donor.GetBloodType(),
		donor.GetIsOrganDonor(),
		donor.GetDonationDate(),
		donor.GetBirthDate(),
	)
	if err != nil {
		logger.Error("Error trying to insert donor in database", err)
		return err
	}
	fmt.Println(result)
	
	return nil
}


// From Spring Boot project
// Optional<Doador> findByIdDoador(Integer idDoador);
//     Optional<Doador> findByEmail(String email);
//     Optional<Doador> findBySenha(String senha);
//     Optional<Doador> findByEmailAndSenha(String email, String senha);
//     String findByNome(String nome);