package shibeentities

type ErrorResponse struct {
	Success bool `json:"success"`
	ErrorCode int64 `json:"code"`
	Message string `json:"message"`
}
