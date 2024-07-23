package model

import (
	pb "github.com/ssonit/common/protos/order"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersCreation struct {
	ID         primitive.ObjectID      `bson:"_id,omitempty"`
	CustomerId string                  `bson:"customerId,omitempty"`
	Status     string                  `bson:"status,omitempty"`
	Items      []*pb.ItemsWithQuantity `bson:"items,omitempty"`
}

type OrderItem struct {
	ID         primitive.ObjectID      `bson:"_id,omitempty"`
	CustomerId string                  `bson:"customerId,omitempty"`
	Status     string                  `bson:"status,omitempty"`
	Items      []*pb.ItemsWithQuantity `bson:"items,omitempty"`
}
