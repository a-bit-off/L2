package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	result := <-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
	)

	if result != nil {
		t.Errorf("Expected 'ch1', but got %v", result)
	}
}
