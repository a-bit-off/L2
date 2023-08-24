package storage

import "fmt"

type Event struct {
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

type Storage struct {
	events map[int]Event
}

func NewEvent(userID int, date string) *Event {
	return &Event{
		UserID: userID,
		Date:   date,
	}
}

func NewStorage() *Storage {
	return &Storage{
		events: make(map[int]Event),
	}
}

func (s *Storage) Add(event *Event) error {
	if _, ok := s.events[event.UserID]; !ok {
		s.events[event.UserID] = *event
		return nil
	}
	return fmt.Errorf("user id not unique")
}

func (s *Storage) Update(event *Event) error {
	if _, ok := s.events[event.UserID]; ok {
		s.events[event.UserID] = *event
		return nil
	}
	return fmt.Errorf("user id not faund")
}

func (s *Storage) Delete(userID int) error {
	if _, ok := s.events[userID]; ok {
		delete(s.events, userID)
		return nil
	}
	return fmt.Errorf("user id not faund")
}
