package structure

type Band struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Debut_Year int   `json:"debut_year"`
}