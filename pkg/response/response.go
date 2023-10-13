package response

import "github.com/EdwardKerckhof/gohtmx/pkg/request"

type successResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type paginatedResponse struct {
	Total   int64 `json:"total"`
	HasMore bool  `json:"hasMore"`
	First   bool  `json:"first"`
	Last    bool  `json:"last"`
	successResponse
}

type errorResponse struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}

func Success(data interface{}) successResponse {
	return successResponse{
		Data:    data,
		Success: true,
	}
}

func Paginated(data interface{}, total int64, req request.PaginationRequest) paginatedResponse {
	return paginatedResponse{
		Total:           total,
		HasMore:         int64(req.Page*req.Size) < total,
		First:           req.Page == 1,
		Last:            int64(req.Page*req.Size) >= total,
		successResponse: Success(data),
	}
}

func Error(err error) errorResponse {
	return errorResponse{
		Success:      false,
		ErrorMessage: err.Error(),
	}
}
