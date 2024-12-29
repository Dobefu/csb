package structs

type Translation struct {
	Source      string `json:"source"`
	Translation string `json:"translation"`
	Category    string `json:"category"`
}
