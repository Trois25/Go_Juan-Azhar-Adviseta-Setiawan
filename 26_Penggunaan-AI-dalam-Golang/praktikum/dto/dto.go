package dto

type RequestData struct {
	Budget  string `json:"budget"`
	Brand   string `json:"brand"`
	Purpose string `json:"purpose"`
}

type Response struct {
	Status         string
	Recommendation string
}
