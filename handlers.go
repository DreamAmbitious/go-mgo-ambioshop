package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"./models/Order"
	"./models/Product"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome <<User>> !")
}

func ProductIndex(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	var products []Product.Product
	c := session.DB("shop").C("products")
	err := c.Find(bson.M{}).All(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed get all books: ", err)
		return
	}

	// Marshal provided interface into JSON structure
	pj, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", pj)
}

func OrderIndex(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	var orders []Order.Order
	c := session.DB("shop").C("order")
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

func ProductShow(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	var product Product.Product
	vars := mux.Vars(r)
	err := session.DB("shop").C("products").
		Find(bson.M{"_id": bson.ObjectIdHex(vars["Id"])}).
		One(&product)
	if err != nil {
		switch err {
		default:
			w.WriteHeader(http.StatusInternalServerError)
		case mgo.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		}
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	pj, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", pj)
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	// Stub an product to be populated from the body
	p := Product.Product{}
	Product.EnsureIndex(session)
	// Populate the product data
	json.NewDecoder(r.Body).Decode(&p)
	log.Println(&p)
	// Add an Id
	p.Id = bson.NewObjectId()
	// Write the product to mongo
	err := session.DB("shop").C("products").Insert(p)
	if err != nil {
		log.Println("I should get here")
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Product with this Name already exists")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed insert product: ", err)
		return
	}

	// Marshal provided interface into JSON structure
	pj, _ := json.Marshal(p)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", pj)
}

func ProductDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	session := getSession()
	defer session.Close()
	err := session.DB("shop").C("products").Remove(bson.M{"_id": bson.ObjectIdHex(vars["Id"])})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		switch err {
		default:
			w.WriteHeader(http.StatusInternalServerError)
		case mgo.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		}
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func OrderCreate(w http.ResponseWriter, r *http.Request) {
	session := getSession()
	defer session.Close()
	o := Order.Order{}
	o.Created = time.Now()
	Product.EnsureIndex(session)
	json.NewDecoder(r.Body).Decode(&o)
	// Add an Id
	o.Id = bson.NewObjectId()
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
