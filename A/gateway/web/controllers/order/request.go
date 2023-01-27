package order

import (
	order2 "gateway/web/controllers/order/order"
	"github.com/golang/protobuf/proto"
)

// Request Very basic validation not even checking numerics...
type Request struct {
	OrderId int32  `json:"order_id" validate:"required"`
	Price   int32  `json:"price" validate:"required"`
	Title   string `json:"title" validate:"required"`
}

func (r *Request) Serialize() ([]byte, error) {
	order := &order2.Order{
		OrderId: r.OrderId,
		Price:   r.OrderId,
		Title:   r.Title,
	}

	return proto.Marshal(order)
}
