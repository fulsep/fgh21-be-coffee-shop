package models

type Products struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	UserId      *int   `json:"userId" db:"user_id"`
}

type JoinProducts struct {
	Id 				int `json:"id"`
	Image 			string `json:"image"`
	Title 			string `json:"title"`
	Price 			int `json:"price"`
	Description 	string `json:"description"`
	ProductSizes 	string `json:"productSizes" db:"product_sizes"`
	OrderType 		string `json:"orderType" db:"order_type"`
	Stock 			int `json:"stock"`
}