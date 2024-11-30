package response

import "time"

type Base struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Timestamp  time.Time   `json:"timestamp"`
	Data       interface{} `json:"data"`
}

type Pagination struct {
	Total       int `json:"total"`   // Total Data
	PerPage     int `json:"perPage"` // Total Data per Page
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"lastPage"`
}
