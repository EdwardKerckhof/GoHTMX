package dto

import "github.com/EdwardKerckhof/gohtmx/pkg/request"

type FindAllRequest struct {
	request.PaginationRequest
	Sort string `form:"sort,default=id" binding:"omitempty,oneof=id username createdAt"`
}
