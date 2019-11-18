package main

import (
	"errors"
	"log"

	"gopkg.in/zabawaba99/firego.v1"
)

var fbUrl = "https://quiet-2c963.firebaseio.com/"

type loc struct {
	lat  float64 `json:"lat"`
	long float64 `json:"long"`
}

type event struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location loc    `json:"loc"`
}

func (p *event) getEvent(s string) (interface{}, error) {
	var f = firego.New(fbUrl+"/event/"+s, nil)
	var v map[string]interface{}
	//f.Shallow(true)

	if err := f.Value(&v); err != nil {
		log.Fatal(err)
		return nil, err
	}

	//fmt.Printf("%s\n\n", v)

	keys := make([]interface{}, 0, len(v))
	for _, value := range v {
		keys = append(keys, value)
	}
	//fmt.Printf("%s\n\n", keys)
	//q := []event{}
	return keys, nil
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

func getEvents() (interface{}, error) {
	//fmt.Println("hi")
	var f = firego.New(fbUrl+"/events", nil)
	var v map[string]interface{}
	//f.Shallow(true)

	if err := f.Value(&v); err != nil {
		log.Fatal(err)
		return nil, err
	}

	//fmt.Printf("%s\n\n", v)

	keys := make([]interface{}, 0, len(v))
	for _, value := range v {
		keys = append(keys, value)
	}
	//fmt.Printf("%s\n\n", keys)
	//q := []event{}
	return keys, nil
}
