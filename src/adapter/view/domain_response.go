package view

import (
	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller/dto/response"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
)

func DonorResponse(donorDomain entity.DonorDomainInterface) response.DonorResponse {
	return response.DonorResponse{
		DonorId:      donorDomain.GetDonorId(),
		Name:         donorDomain.GetName(),
		// Cpf:          donorDomain.GetCpf(),
		// Email:        donorDomain.GetEmail(),
		// Gender:       donorDomain.GetGender(),
		// BloodType:    donorDomain.GetBloodType(),
		// IsOrganDonor: donorDomain.GetIsOrganDonor(),
		// DonationDate: donorDomain.GetDonationDate(),
	}
}

func PostResponse(postDomain entity.PostDomainInterface) response.PostResponse {
	return response.PostResponse{
		PostId:      postDomain.GetPostId(),
		Description: postDomain.GetDescription(),
		PostDate:    postDomain.GetPostDate(),
		Receiver:    postDomain.GetReceiver(),
	}
}
