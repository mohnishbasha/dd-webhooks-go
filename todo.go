package main

import "time"

type Webhook struct {
	Id        int       `json:"id"`
        EventId   int       `json:"eventid"`
        OrgId     int       `json:"orgid"`
        Title     string    `json:"title"`
        EventType string    `json:"eventtype"`
        EventBody string    `json:"eventbody"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Webhooks []Webhook
