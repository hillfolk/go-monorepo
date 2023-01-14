package orders

import (
	"fmt"
	"time"
)

type Order struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	ProductId int64     `json:"productId"`
	Quantity  int64     `json:"quantity"`
	Price     float64   `json:"price"`
	OrderTime time.Time `json:"orderTime"`
}

func NewOrder(id int64, userId int64, productId int64, quantity int64, price float64) *Order {
	return &Order{
		Id:        id,
		UserId:    userId,
		ProductId: productId,
		Quantity:  quantity,
		Price:     price,
		OrderTime: time.Now(),
	}
}

func (o *Order) ToString() string {
	return fmt.Sprintf("Order: %d, %d, %d, %d, %f, %s", o.Id, o.UserId, o.ProductId, o.Quantity, o.Price, o.OrderTime.Format("2006-01-02 15:04:05"))
}
