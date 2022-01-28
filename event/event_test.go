package event

import (
	"testing"
)

// Test som tester PutChickenBoat
func TestPutChickenBoat(t *testing.T) {
	wanted := "Kyllingen er i båten"
	got := PutChickenBoat
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q. Enten er det ingen i båten, for mange i båten eller Mennesket mangler", got, wanted)
	} else {
		t.Log("Kyllingen er i båten")
	}
}
