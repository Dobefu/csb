package structs

type QueryType int

const (
	QueryTypeSelect QueryType = iota
	QueryTypeInsert
	QueryTypeUpsert
	QueryTypeUpdate
	QueryTypeDelete
)

type Query struct {
	Type   QueryType
	Fields []string
	Where  []QueryWhere
}
