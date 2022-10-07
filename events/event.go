package main

import "time"

type ConflictMessage struct {
	To        string  `json:"to"`
	Conflicts []Event `json:"conflicts"`
}

type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Organizer string    `json:"organizer"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}
