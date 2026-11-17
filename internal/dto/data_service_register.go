package dto

type DataServiceRegister struct {
	PartName string `json:"part_name" binding:"required" example:"Oli Mesin"`
	Amount   int16  `json:"amount" binding:"required" example:"1"`
}
