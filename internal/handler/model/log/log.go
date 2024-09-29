package log

import "time"

type Log struct {
	ID      int64  `json:"id"`
	TraceID string `json:"trace_id"`
	Module  string `json:"module"`

	Time  time.Time `json:"time"`
	Level string    `json:"level"`

	Message string `json:"message"`
}

type Filter struct {
	TraceID  *string `json:"trace_id"`
	ModuleID *int64  `json:"module_id"`

	Level   *string `json:"level"`
	Message string  `json:"message"`

	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`

	CountOnPage int64 `json:"count_on_page" binding:"required"`
	Page        int64 `json:"page" binding:"required"`
}

type GetResp struct {
	Logs []Log `json:"logs"`

	PagesCount int64 `json:"pages_count"`
}
