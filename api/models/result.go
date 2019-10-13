package models

type Result struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
}

func NewResult() *Result {
	return resultPool.Get().(*Result)
}

func (r *Result) Release() {
	r.Status = 200
	r.Data = nil
	r.Code = ""
	r.Message = ""
	resultPool.Put(r)
}
