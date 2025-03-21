package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52 but got %v", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts as the first card but received %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs as the last card but received %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck of length 52 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
