package samples

import (
	"time"
)

//Post ...
type Post struct {
	ID        int        `gorm:"primary_key;auto_increment;type:int(11);not null;column:id" json:"id"`
	Name      string     `gorm:"type:varchar(191);unique;not null;unique_index:name_UNIQUE;column:name" json:"name"`
	Content   string     `gorm:"type:text;column:content" json:"content,omitempty"`
	CreatedAt time.Time  `gorm:"type:datetime;column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:datetime;column:updated_at" json:"updated_at,omitempty"`
	Category  *Category  `gorm:"foreignkey:ID" json:"category,omitempty"`
}

//TableName ...
func (Post) TableName() string {
	return "post"
}
