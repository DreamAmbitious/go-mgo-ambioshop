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
}
