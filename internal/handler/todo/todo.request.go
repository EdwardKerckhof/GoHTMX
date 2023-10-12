package todo

import "github.com/EdwardKerckhof/gohtmx/pkg/request"

type findAllRequest struct {
	request.PaginationRequest
	Sort string `form:"sort,default=id" binding:"omitempty,oneof=id title createdAt"`
}

type createRequest struct {
	Title string `json:"title" form:"title" binding:"required,min=1,max=255"`
}

type updateRequest struct {
	Title     string `json:"title" form:"title" binding:"required,min=1,max=255"`
	Completed bool   `json:"completed" form:"completed" binding:"omitempty"`
}
