package state

import "testing"

func TestTsViewState(t *testing.T) {
	wanted := "[kylling rev korn hs ---\\ \\__/ _________________/---]"
	state := TsViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
