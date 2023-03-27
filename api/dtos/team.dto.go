package dtos

type TeamRequest struct {
	Name  string `json:"name"`
	Pages int64  `json:"pages"`
	Limit *int64 `json:"limit,omitempty"`
}
