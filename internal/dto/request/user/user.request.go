package user

import "github.com/EdwardKerckhof/gohtmx/internal/dto/request"

type FindAllRequest struct {
	request.PaginationRequest
	Sort string `form:"sort,default=id" binding:"omitempty,oneof=id username createdAt"`
}
