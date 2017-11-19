package main

import "net/http"

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
		Index,
	},
	Route{
		"ProductIndex",
		"GET",
		"/products",
		ProductIndex,
	},
	Route{
		"ProductCreate",
		"POST",
		"/products",
		ProductCreate,
	},
	Route{
		"ProductDelete",
		"DELETE",
		"/products/{Id}",
		ProductDelete,
	},
	Route{
		"ProductShow",
		"GET",
		"/products/{Id}",
		ProductShow,
	},

	Route{
		"OrderIndex",
		"GET",
		"/orders",
		OrderIndex,
	},
	Route{
		"OrderCreate",
		"POST",
		"/orders",
		OrderCreate,
	},
	Route{
		"OrderDelete",
		"DELETE",
		"/products/{Id}",
		ProductDelete,
	},
	Route{
		"OrderShow",
		"GET",
		"/products/{Id}",
		ProductShow,
	},
}
