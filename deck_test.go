package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards, got %v", len(d))
	}
}

func TestFirstElementOfDeck(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, got %v", d[0])
	}
}

func TestLastElementOfDeck(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
