package dto

import "go_bengkel/internal/models"

type DataServiceRegister struct {
	PartName string `json:"part_name" binding:"required"`
	Amount   int16  `json:"amount" binding:"required"`
}

func ToDataServiceRegister(dataService models.DataService) DataServiceRegister {
	return DataServiceRegister{
		PartName: dataService.PartName,
		Amount:   dataService.Amount,
	}
}
