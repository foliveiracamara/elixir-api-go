package controller

import (
	"net/http"

	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller/dto/request"
	"github.com/foliveiracamara/elixir-api-go/src/adapter/view"
	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	"github.com/foliveiracamara/elixir-api-go/src/config/validation"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
	"github.com/foliveiracamara/elixir-api-go/src/domain/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	DonorDomainInterface entity.DonorDomainInterface
)

type DonorController struct {
	service service.DonorDomainService
}

func NewDonorController(svc service.DonorDomainService)  DonorController {
	return DonorController{
		service: svc,
	}
}

// Create a new blood donor
func (d *DonorController) CreateDonor(ctx *gin.Context) {
	var donorRequest request.DonorRequest
	if err := ctx.ShouldBindJSON(&donorRequest); err != nil {
		logger.Error(
			"Error trying to validate donor info", err,
			zap.String("journey", "createDonorController"),
		)
		appErr := validation.ValidateDonorError(err)
		ctx.JSON(appErr.Code, appErr)
		return
	}

	newDonor := entity.NewDonorDomain(
		donorRequest.DonorId,
		donorRequest.Name,
		donorRequest.Cpf,
		donorRequest.Email,
		donorRequest.Password,
		donorRequest.Gender,
		donorRequest.BloodType,
		donorRequest.DonationDate,
		donorRequest.BirthDate,
		donorRequest.IsOrganDonor,
	)

	if err := d.service.CreateDonor(newDonor); err != nil {
		ctx.JSON(err.Code, err)
	}
	
	logger.Info("Donor created successfully",
		zap.String("journey", "createDonorController"),
	)
	
	ctx.JSON(http.StatusOK, view.DonorResponse(newDonor))
}

func (d *DonorController) ListDonors(ctx *gin.Context) {

}

func (d *DonorController) GetDonorById(ctx *gin.Context) error {
	// id := ctx.Param("id")
	// for _, donor := range entity.Donors {
	// 	if donor.DonorId == id {
	// 		return ctx.JSON(http.StatusOK, donor)
	// 	}
	// }
	// return ctx.JSON(http.StatusBadRequest, nil)
	return nil
}

func (d *DonorController) UpdateDonor(ctx *gin.Context) error {
	// id := ctx.Param("id")
	// for i := range entity.Donors {
	// 	if entity.Donors[i].DonorId == id {
	// 		entity.Donors[i].IsOrganDonor = true
	// 		return ctx.JSON(http.StatusOK, entity.Donors[i])
	// 	}
	// }
	// return ctx.JSON(http.StatusNotFound, nil)
	return nil
}

func (d *DonorController) DeleteDonor(ctx *gin.Context) error {
	// id := ctx.Param("id")
	//
	//	for i := range entity.Donors {
	//		if entity.Donors[i].DonorId == id {
	//			entity.Donors = append(entity.Donors[:i], entity.Donors[i+1:]...)
	//			return ctx.JSON(http.StatusOK, nil)
	//		}
	//	}
	//
	// return ctx.JSON(http.StatusNotFound, nil)
	return nil
}

// func LoginDonor(ctx echo.Context) error {
// 	email := ctx.Param("email")
// 	password := ctx.Param("password")
// 	for i := range entity.Donors {
// 		if entity.Donors[i].Email == email && entity.Donors[i].Password == password {
// 			return ctx.JSON(http.StatusOK, nil)
// 		}
// 	}
// 	return ctx.JSON(http.StatusNotFound, nil)
// }
