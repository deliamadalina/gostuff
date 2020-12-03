package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52, got  %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Expected Ten of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected first card of Ace of Spades but got %v", len(loadedDeck))
	}

}
