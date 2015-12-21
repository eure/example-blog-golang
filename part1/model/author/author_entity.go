package author

import (
	"fmt"
)

// Entity is struct for author
type Entity struct {
	Email string `response:"email" xorm:"pk not null unique"`
	Name  string `response:"name" xorm:"not null unique"`
}

const (
	tableName = "authors"
	pkName    = "email"
)

// TableName define table name for Xorm
func (e Entity) TableName() string {
	return tableName
}

func (e Entity) String() string {
	return fmt.Sprintf("{email: %s, name: %s}", e.Email, e.Name)
}
