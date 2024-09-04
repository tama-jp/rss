package request

type TimeRequest struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}
