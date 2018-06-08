package models

// Consecutive ....
type Consecutive struct {
	tableName struct{} `sql:"consecutive"`

	ID   string `sql:"type:text, pk"`
	Type string `sql:"type:text"` //FK

}
