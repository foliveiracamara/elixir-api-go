package controller

import (
	"net/http"

	"github.com/foliveiracamara/elixir-api-go/domain/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ListDonors(c echo.Context) error {
	if entity.Donors == nil{
		return c.String(http.StatusNotFound, "No user found")
	}
	return c.JSON(http.StatusOK, entity.Donors)
}

// func GetDonor(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("DonorId"))
// 	for i, _ := range entity.Donors{
// 		if entity.Donor.DonorId == id {
// 			return c.JSON(http.StatusOK, entity.Donor)
// 		}
// 	}
// 	return c.JSON(http.StatusBadRequest, nil)
// }

func CreateDonor(c echo.Context) error {
	newDonor := entity.Donor{}	
	err := c.Bind(&newDonor)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	newDonor.DonorId = generateUUID()
	entity.Donors = append(entity.Donors, newDonor)
	return c.JSON(http.StatusCreated, entity.Donors)
}

func generateUUID() string {
	uuid := uuid.New()
  return uuid.String() 
}


