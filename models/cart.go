package models

import "time"

type Cart struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int                  `json:"product_id"`
	Product       ProductResponse      `json:"product"`
	UserID        int                  `json:"user_id"`
	User          UsersProfileResponse `json:"user"`
	TransactionID *int                 `json:"transaction_id"`
	Transaction   TransactionResponse  `json:"transaction"`
	Qty           int                  `json:"qty"`
	CreatedAt     time.Time            `json:"-"`
	UpdatedAt     time.Time            `json:"updated_at"`
	SubAmount     int                  `json:"sub_amount" gorm:"type:int"`
}

type CartResponse struct {
	ID            int                 `json:"id"`
	ProductID     int                 `json:"product_id"`
	Product       Product             `json:"product"`
	TransactionID int                 `json:"-"`
	Transaction   TransactionResponse `json:"transaction" gorm:"foreignKey:TransactionID"`
	SubAmount     int                 `json:"sub_amount"`
}

func (CartResponse) TableName() string {
	return "carts"
}
