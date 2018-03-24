package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DreamAmbitious/go-mgo-ambioshop.git/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func OrderCreate(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	o := models.Order{}
	o.Created = time.Now()
	json.NewDecoder(r.Body).Decode(&o)
	// Add an Id
	o.Id = bson.NewObjectId()
	log.Println(&o)
	// Write the product to mongo
	err := session.DB("shop").C("orders").Insert(&o)

	if err != nil {
		log.Println("I should get here")
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unable to take your order at this time")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed insert order: %v", err)
		return
	}
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
}

func OrderIndex(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	var orders []models.Order
	c := session.DB("shop").C("orders")
	err := c.Find(bson.M{}).All(&orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed get all books: ", err)
		return
	}

	// Marshal provided interface into JSON structure
	pj, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", pj)
}
