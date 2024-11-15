package models

type TotalSales struct {
	Amount float64 `json:"total_sales"`
}

type MostPopularItem struct {
	Item     string `json:"most_popular_item"`
	Quantity int    `json:"quantity"`
}
