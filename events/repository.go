package main

type Repository interface {
	AddEvent(event Event) error
	//ListEvents(organizer string) ([]Event, error)
}
