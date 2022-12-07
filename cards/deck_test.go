package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expecting 16")
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected ace of spades")
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected four of clubs")
	}
}

func TestSavetodeckandnewdeckfromfile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckfromFile("_decktesting")
	if len(ld) != 16 {
		t.Errorf("Expecting 16")
	}
	os.Remove("_decktesting")
}
