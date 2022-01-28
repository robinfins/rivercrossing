package state

import (
	"strconv"
)

//Tilstander som skal kunne endres gjennom Event
var ChickenWest = true
var FoxWest = true
var CornWest = true
var HsWest = true

var ChickenBoat = false
var FoxBoat = false
var CornBoat = false
var HsBoat = false

var ChickenEast = false
var FoxEast = false
var CornEast = false
var HsEast = false

//Tilstanden til hvem som er på båten
func StateBoat() (FinalState string) {

	FinalState = "Sts Chicken Boat:" + strconv.FormatBool(ChickenBoat) + " | " + "Sts Fox Boat:" + strconv.FormatBool(FoxBoat) + " | " + "Sts Corn Boat:" + strconv.FormatBool(CornBoat) + " | " + "Sts Hs Boat:" + strconv.FormatBool(HsBoat)

	return
}

//Tilstanden til hvem som er på land vest
func StateLandV() (FinalState string) {

	FinalState = "Sts Chicken West:" + strconv.FormatBool(ChickenWest) + " | " + "Sts Fox West:" + strconv.FormatBool(FoxWest) + " | " + "Sts Corn West:" + strconv.FormatBool(CornWest) + " | " + "Sts Hs West:" + strconv.FormatBool(HsWest)

	return
}

//Tilstand til hvem som er på land øst
func StateLandE() (FinalState string) {

	FinalState = "Sts Chicken East:" + strconv.FormatBool(ChickenEast) + " | " + "Sts Fox East:" + strconv.FormatBool(FoxEast) + " | " + "Sts Corn East:" + strconv.FormatBool(CornEast) + " | " + "Sts Hs East:" + strconv.FormatBool(HsEast)

	return
}

//Alle tilstandene lagt sammen
func FinalState() (FinalState string) {
	FinalState = StateBoat() + StateLandE() + StateLandV()
	return
}

//Funksjon for å vise Final State med ViewState
func ViewState() string {
	return FinalState()
}

//Funksjon for tViewState
func TsViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return "[kylling rev korn hs ---\\ \\__/ _________________/---]"
}
