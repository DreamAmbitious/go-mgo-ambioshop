package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Customerid string        `json:"customerid" bson:"customerid"`
	Totalprice float64       `json:"totalprice" bson:"totalprice"`
	Created    time.Time     `json:"created" bson:"created"`
	Modified   time.Time     `json:"modified" bson:"modified"`
	Status     string        `json:"status" bson:"status"`
	Orderitems []*Orderitem
}

type Orderitem struct {
	Productid string `json:"productid" bson:"productid"`
	Quantity  string `json:"quantity" bson:"quantity"`
}

func OrderIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("shop").C("orders")

	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
