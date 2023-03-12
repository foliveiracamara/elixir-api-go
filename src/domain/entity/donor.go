package entity

type Donor struct {
	DonorId       string `json:"donorId"`
	Name          string `json:"name"`
	Cpf           string `json:"cpf"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Gender        string `json:"gender,"`
	BloodType     string `json:"bloodType"`
	IsOrganDonor  bool   `json:"isOrganDonor"`
	PicturePerson string `json:"picturePerson"`
	DonationDate  string `json:"donationDate"`
}

type donors []Donor

var Donors donors
