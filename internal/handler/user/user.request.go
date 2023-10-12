package user

import "github.com/EdwardKerckhof/gohtmx/pkg/request"

type findAllRequest struct {
	request.PaginationRequest
	Sort string `form:"sort,default=id" binding:"omitempty,oneof=id username createdAt"`
}
