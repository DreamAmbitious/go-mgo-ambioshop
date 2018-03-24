package main

import (
	"net/http"

	"github.com/DreamAmbitious/go-mgo-ambioshop.git/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"ProductIndex",
		"GET",
		"/products",
		handlers.ProductIndex,
	},
	Route{
		"ProductCreate",
		"POST",
		"/products",
		handlers.ProductCreate,
	},
	Route{
		"ProductDelete",
		"DELETE",
		"/products/{Id}",
		handlers.ProductDelete,
	},
	Route{
		"ProductShow",
		"GET",
		"/products/{Id}",
		handlers.ProductShow,
	},

	Route{
		"OrderIndex",
		"GET",
		"/orders",
		handlers.OrderIndex,
	},
	Route{
		"OrderCreate",
		"POST",
		"/orders",
		handlers.OrderCreate,
	},
	/*			Route{
					"OrderDelete",
					"DELETE",
					"/products/{Id}",
					handlers.ProductDelete,
				},
				Route{
					"OrderShow",
					"GET",
					"/products/{Id}",
					handlers.ProductShow,
				},*/
}
