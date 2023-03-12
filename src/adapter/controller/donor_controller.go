package controller

import (
	"net/http"

	"github.com/foliveiracamara/elixir-api-go/src/adapter/utils"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
	"github.com/labstack/echo/v4"
)

func ListDonors(ctx echo.Context) error {
	if entity.Donors == nil {
		return ctx.String(http.StatusNotFound, "No user found")
	}
	return ctx.JSON(http.StatusOK, entity.Donors)
}

func GetDonorById(ctx echo.Context) error {
	id := ctx.Param("id")
	for _, donor := range entity.Donors {
		if donor.DonorId == id {
			return ctx.JSON(http.StatusOK, donor)
		}
	}
	return ctx.JSON(http.StatusBadRequest, nil)
}

func CreateDonor(ctx echo.Context) error {
	newDonor := entity.Donor{}
	err := ctx.Bind(&newDonor)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	newDonor.DonorId = utils.GenerateUUID()
	entity.Donors = append(entity.Donors, newDonor)
	return ctx.JSON(http.StatusCreated, entity.Donors)
}

func UpdateDonor(ctx echo.Context) error {
	id := ctx.Param("id")
	for i := range entity.Donors {
		if entity.Donors[i].DonorId == id {
			entity.Donors[i].IsOrganDonor = true
			return ctx.JSON(http.StatusOK, entity.Donors[i])
		}
	}
	return ctx.JSON(http.StatusNotFound, nil)
}

func DeleteDonor(ctx echo.Context) error {
	id := ctx.Param("id")
	for i := range entity.Donors {
		if entity.Donors[i].DonorId == id {
			entity.Donors = append(entity.Donors[:i], entity.Donors[i+1:]...)
			return ctx.JSON(http.StatusOK, nil)
		}
	}
	return ctx.JSON(http.StatusNotFound, nil)
}

func LoginDonor(ctx echo.Context) error {
	email := ctx.Param("email")
	password := ctx.Param("password")
	for i := range entity.Donors {
		if entity.Donors[i].Email == email && entity.Donors[i].Password == password {
			return ctx.JSON(http.StatusOK, nil)
		}
	}
	return ctx.JSON(http.StatusNotFound, nil)
}


