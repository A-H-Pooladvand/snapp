package models

import pb "order/proto"

type Order struct {
	Model
	ID    int32  `json:"id" gorm:"primarykey"`
	Title string `json:"title"`
	Price int32  `json:"price"`
}

func (o *Order) Persist() *Model {
	return DB().Create(o)
}

func NewOrderFromProto(o *pb.Order) *Order {
	return &Order{
		ID:    o.OrderId,
		Title: o.Title,
		Price: o.Price,
	}
}
