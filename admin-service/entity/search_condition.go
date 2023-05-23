package entity

type Comparison int

const (
	Equal Comparison = iota + 1
	NotEqual
	GreaterThan
	GreaterThanOrEqual
	LessThan
	LessThanOrEqual
	Like
)

type SearchCondition struct {
	Field     string     `json:"field"`
	FieldName string     `json:"fieldName"`
	FieldType string     `json:"fieldType"`
	Condition Comparison `json:"condition"`
	Value     string     `json:"value"`
}
