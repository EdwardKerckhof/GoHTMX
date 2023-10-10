package request

type PaginationRequest struct {
	Page  int32  `form:"page,default=1" binding:"omitempty,min=1"`
	Size  int32  `form:"size,default=50" binding:"omitempty,min=1,max=100"`
	Sort  string `form:"sort,default=id" binding:"omitempty,oneof=id createAt lastLoginAt"`
	Order string `form:"order,default=asc" binding:"omitempty,oneof=asc desc"`
}
