package samples

import (
	"time"
)

//Category ...
type Category struct {
	ID          int        `gorm:"primary_key;auto_increment;type:int(11);not null;column:id" json:"id"`
	Name        string     `gorm:"type:varchar(191);unique;not null;unique_index:name_UNIQUE;column:name" json:"name"`
	Description string     `gorm:"type:varchar(191);column:description" json:"description,omitempty"`
	CreatedAt   time.Time  `gorm:"type:datetime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"type:datetime;column:updated_at" json:"updated_at,omitempty"`
}

//TableName ...
func (Category) TableName() string {
	return "category"
}
