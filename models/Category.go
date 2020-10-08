package models

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null;unique" json:"name"`

	Articles []Article `gorm:"many2many:article_category_relation;"`
}
