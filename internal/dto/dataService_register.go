package dto

import "go_bengkel/internal/models"

type DataServiceRegister struct {
	PartName string `json:"part_name" binding:"required" example:"Oli Mesin"`
	Amount   int16  `json:"amount" binding:"required" example:"1"`
}

func ToDataServiceRegister(dataService models.DataService) DataServiceRegister {
	return DataServiceRegister{
		PartName: dataService.PartName,
		Amount:   dataService.Amount,
	}
}
