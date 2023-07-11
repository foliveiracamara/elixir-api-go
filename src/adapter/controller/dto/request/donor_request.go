package request

type DonorRequest struct {
	DonorId      string `json:"donor_id"`
	Name         string `json:"name"      binding:"required,min=4,max=60"`
	Cpf          string `json:"cpf"       binding:"required,len=11"`
	Email        string `json:"email"     binding:"required,email,min=8,max=50"`
	Password     string `json:"password"  binding:"required,min=10,max=150"`
	BloodType    string `json:"bloodType" binding:"required,len=2"`
	Gender       string `json:"gender"`
	IsOrganDonor bool   `json:"isOrganDonor"`
	DonationDate string `json:"donationDate"`
	BirthDate    string `json:"birthDate" binding:"required"`
}
