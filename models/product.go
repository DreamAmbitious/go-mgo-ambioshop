package models

import mgo "gopkg.in/mgo.v2"

func EnsureIndex(s *mgo.Session) {
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
