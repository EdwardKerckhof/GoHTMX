package request

import (
	"fmt"

	"github.com/google/uuid"
)

type IDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type PaginationRequest struct {
	Page  int32  `form:"page,default=1" binding:"omitempty,min=1"`
	Size  int32  `form:"size,default=50" binding:"omitempty,min=1,max=100"`
	Order string `form:"order,default=asc" binding:"omitempty,oneof=asc desc"`
}

func (r *IDRequest) ParseID() (uuid.UUID, error) {
	uuid, err := uuid.Parse(r.ID)
	if err != nil {
		return uuid, fmt.Errorf("invalid uuid: %w", err)
	}
	return uuid, nil
}
