package storage

import (
	"encoding/json"
	"fmt"
	"time"
)

// map:
// key: int (user id)
// value: map
//		key: int (event id)
//		value: string (date)

type Storage struct {
	Users map[int]map[int]time.Time `json:"users"`
}

type DateInfo struct {
	EventID int
	Date    time.Time
}

func NewStorage() *Storage {
	return &Storage{
		Users: make(map[int]map[int]time.Time),
	}
}

func (s *Storage) Add(userID, eventID int, date string) error {
	if _, ok := s.Users[userID]; !ok {
		s.Users[userID] = make(map[int]time.Time)
	}

	if _, ok := s.Users[userID][eventID]; !ok {
		dateTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			return fmt.Errorf("date error")
		}
		s.Users[userID][eventID] = dateTime
		return nil
	}
	return fmt.Errorf("event id not unique")
}

func (s *Storage) Delete(userID, eventID int) error {
	if _, ok := s.Users[userID]; !ok {
		return fmt.Errorf("user id not found")
	}
	if _, ok := s.Users[userID][eventID]; !ok {
		return fmt.Errorf("event id not found")
	}

	delete(s.Users[userID], eventID)

	return nil
}

func (s *Storage) Update(userID, eventID int, date string) error {
	if _, ok := s.Users[userID]; !ok {
		return fmt.Errorf("user id not found")
	}
	if _, ok := s.Users[userID][eventID]; !ok {
		return fmt.Errorf("event id not found")
	}

	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("date error")
	}
	s.Users[userID][eventID] = dateTime

	return nil
}

func (s *Storage) ForDay(userID, eventDay int) (string, error) {
	if _, ok := s.Users[userID]; !ok {
		return "", fmt.Errorf("user id not found")
	}

	var dates []DateInfo
	for k, v := range s.Users[userID] {
		_, _, d := v.Date()
		if d == eventDay {
			dates = append(dates, DateInfo{EventID: k, Date: v})
		}
	}

	byt, err := json.Marshal(&dates)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}

func (s *Storage) ForWeek(userID, eventWeek int) (string, error) {
	if _, ok := s.Users[userID]; !ok {
		return "", fmt.Errorf("user id not found")
	}

	var dates []DateInfo
	for k, v := range s.Users[userID] {
		_, w := v.ISOWeek()
		if w == eventWeek {
			dates = append(dates, DateInfo{EventID: k, Date: v})
		}
	}

	byt, err := json.Marshal(&dates)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}

func (s *Storage) ForMonth(userID, eventMonth int) (string, error) {
	if _, ok := s.Users[userID]; !ok {
		return "", fmt.Errorf("user id not found")
	}

	var dates []DateInfo
	for k, v := range s.Users[userID] {
		_, m, _ := v.Date()
		if int(m) == eventMonth {
			dates = append(dates, DateInfo{EventID: k, Date: v})
		}
	}

	byt, err := json.Marshal(&dates)
	if err != nil {
		return "", err
	}

	return string(byt), nil
}
