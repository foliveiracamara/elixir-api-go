package entity

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/foliveiracamara/elixir-api-go/src/adapter/utils"
)

type DonorDomainInterface interface {
	GetDonorId() string
	GetName() string
	GetCpf() string
	GetEmail() string
	GetPassword() string
	GetGender() string
	GetBloodType() string
	GetDonationDate() string
	GetBirthDate() string
	GetIsOrganDonor() bool
	EncryptPassword()
	GenerateId()
}

func NewDonorDomain(
	donorId, 
	name, 
	cpf, 
	email, 
	password, 
	gender, 
	bloodType, 
	donationDate, 
	birthDate string,
	isOrganDonor bool,
) DonorDomainInterface {
	return &donorDomain{
		donorId, 
		name, 
		cpf, 
		email, 
		password, 
		gender, 
		bloodType, 
		donationDate, 
		birthDate,
		isOrganDonor,
	}
}

type donorDomain struct {
	donorId      string
	name         string
	cpf          string
	email        string
	password     string
	gender       string
	bloodType    string
	donationDate string
	birthDate 	 string
	isOrganDonor bool
}

func (d *donorDomain) GetDonorId() string {
	return d.donorId
}

func (d *donorDomain) GetName() string {
	return d.name
}

func (d *donorDomain) GetCpf() string {
	return d.cpf
}

func (d *donorDomain) GetEmail() string {
	return d.email
}

func (d *donorDomain) GetPassword() string {
	return d.password
}

func (d *donorDomain) GetGender() string {
	return d.gender
}

func (d *donorDomain) GetBloodType() string {
	return d.bloodType
}

func (d *donorDomain) GetDonationDate() string {
	return d.donationDate
}

func (d *donorDomain) GetIsOrganDonor() bool {
	return d.isOrganDonor
}

func (d *donorDomain) GetBirthDate() string {
	return d.birthDate
}

func (d *donorDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(d.password))
	d.password = hex.EncodeToString(hash.Sum(nil))
}

func (d *donorDomain) GenerateId() {
	d.donorId = utils.GenerateUUID()
}

