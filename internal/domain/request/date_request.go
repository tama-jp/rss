package request

type ReqDate struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Day     int `json:"day"`
	Weekday int `json:"weekday"`
}
