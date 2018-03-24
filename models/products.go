package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Data struct for Product table
type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Price       float64       `json:"price" bson:"price"`
	Created     time.Time     `json:"created" bson:"created"`
	modified    time.Time     `json:"modified" bson:"modified"`
	Status      string        `json:"status" bson:"status"`
}

func ProductIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("shop").C("products")

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
