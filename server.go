package main

import mgo "gopkg.in/mgo.v2"
import "./models"

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	models.EnsureIndex(s)
	return s
}
