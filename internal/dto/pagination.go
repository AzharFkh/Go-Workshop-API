package dto

import "go_bengkel/internal/utils"

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int64 `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
}

type PaginatedResponse struct {
	Data interface{}    `json:"data"`
	Meta PaginationMeta `json:"metadata"`
}

func ToPaginatedResponse[T any](dtoData interface{}, result utils.PaginatedResult[T]) PaginatedResponse {
	return PaginatedResponse{
		Data: dtoData,
		Meta: PaginationMeta{
			CurrentPage: result.Page,
			PageSize:    result.Limit,
			TotalPages:  result.TotalPages,
			TotalItems:  result.TotalItems,
		},
	}
}
