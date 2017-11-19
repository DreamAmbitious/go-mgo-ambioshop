package main

import mgo "gopkg.in/mgo.v2"
import "./models/Product"
import "./models/Order"

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	Product.EnsureIndex(s)
	Order.EnsureIndex(s)
	return s
}
