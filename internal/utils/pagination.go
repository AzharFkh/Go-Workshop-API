package utils

type PaginatedResult[T any] struct {
	Data       []T
	Page       int
	Limit      int
	TotalPages int64
	TotalItems int64
}
