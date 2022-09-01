package models

type Product struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: int"`
	Description string `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty         int    `json:"-" form:"qty"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type ProductUserResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Qty         int    `json:"qty"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}
