// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameArchitect = "architects"

// Architect mapped from table <architects>
type Architect struct {
	ID   int32  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName Architect's table name
func (*Architect) TableName() string {
	return TableNameArchitect
}
