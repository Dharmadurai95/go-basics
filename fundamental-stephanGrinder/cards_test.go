package main

import "testing"

func TestNewDeck(t *testing.T) {

	newDeck := newCards()
	if len(newDeck) != 16 {
		t.Errorf("Expected 16 card but got %d", len(newCards()))
	}

}
