package models

type Result struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResult() *Result {
	return &Result{
		Status:200,
	}
}
