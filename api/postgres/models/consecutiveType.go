package models

// ConsecutiveType ....
type ConsecutiveType struct {
	tableName struct{} `sql:"consecutive_type"`

	ID   string `sql:"type:text, pk"`
	Name string `sql:"type:text"`
}
