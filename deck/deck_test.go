package main

import (
	"os"
	"testing"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expected deck length 48, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Space, but instead got %s", d[0])
	}
	
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, but instead got %s", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected deck length 48, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
	
}
