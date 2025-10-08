package types

type BandRequest struct {
	Name      string `json:"name"`
	Country   string `json:"country"`
	DebutYear int    `json:"debut_year"`
}

type BandResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	DebutYear int    `json:"debut_year"`
}