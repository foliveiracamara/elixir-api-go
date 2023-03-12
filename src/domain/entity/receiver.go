package entity

type Receiver struct {
	ReceiverId      string `json:"receiverId,omitempty"`
	Name            string `json:"name,omitempty"`
	Cpf             string `json:"cpf,omitempty"`
	BirthdayDate    string `json:"birthdayDate,omitempty"`
	HospitalName    string `json:"hospitalName,omitempty"`
	Cep             string `json:"cep,omitempty"`
	BloodType       string `json:"bloodType,omitempty"`
	Gender          string `json:"gender,omitempty"`
	CellphoneNumber string `json:"cellphoneNumber,omitempty"`
}
