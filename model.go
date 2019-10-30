package main

import (
	"errors"
)

type event struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (p *event) getEvent() error {
	return errors.New("Not implemented")
}

func (p *event) updateEvent() error {
	return errors.New("Not implemented")
}

func (p *event) deleteEvent() error {
	return errors.New("Not implemented")
}

func (p *event) createEvent() error {
	return errors.New("Not implemented")
}

func getEvents() ([]event, error) {
	return nil, errors.New("Not implemented")
}
