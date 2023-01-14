package orders

type Order struct {
	Id        int64   `json:"id"`
	UserId    int64   `json:"userId"`
	ProductId int64   `json:"productId"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}
