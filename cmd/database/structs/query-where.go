package structs

type QueryWhere struct {
	Name     string   `json:"name"`
	Value    any      `json:"value"`
	Operator Operator `json:"operator"`
}
