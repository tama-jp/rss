package response

type DateResponse struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Day     int `json:"day"`
	Weekday int `json:"weekday"`
}
