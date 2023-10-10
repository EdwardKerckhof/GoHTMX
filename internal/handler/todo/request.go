package todo

import "github.com/EdwardKerckhof/gohtmx/pkg/request"

type findByIdRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type findAllRequest struct {
	request.PaginationRequest
}

type createRequest struct {
	Title string `json:"title" binding:"required,min=1,max=255"`
}
