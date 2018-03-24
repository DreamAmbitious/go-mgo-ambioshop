package handlers

import (
	"github.com/DreamAmbitious/go-mgo-ambioshop.git/models"
	mgo "gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	models.ProductIndex(s)
	models.OrderIndex(s)
	return s
}
