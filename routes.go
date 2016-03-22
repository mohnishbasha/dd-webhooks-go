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
		"WebhookIndex",
		"GET",
		"/webhooks",
		WebhookIndex,
	},
	Route{
		"WebhookCreate",
		"POST",
		"/webhooks",
		WebhookCreate,
	},
	Route{
		"WebhookShow",
		"GET",
		"/webhooks/{webhookId}",
		WebhookShow,
	},
}
