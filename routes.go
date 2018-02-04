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
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"SlaShow",
		"GET",
		"/api/v1/sla/{slaId}",
		SlaShow,
	},
	Route{
		"SlaIndex",
		"GET",
		"/api/v1/sla",
		SlaIndex,
	},
	Route{
		"ObjectIndex",
		"GET",
		"/api/v1/objects",
		ObjectIndex,
	},
	Route{
		"ObjectShow",
		"GET",
		"/api/v1/objects/{objId}",
		ObjectShow,
	},
}
