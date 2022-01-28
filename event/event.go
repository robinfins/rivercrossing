package event

import "rivercrossing/state"

//Setter Kylling i båt
func PutChickenBoat() {
	state.ChickenBoat = true
	state.ChickenWest = false
	state.ChickenEast = false
}

//Setter Rev i båt
func PutFoxBoat() {
	state.FoxBoat = true
	state.FoxEast = false
	state.FoxWest = false
}

//Setter Korn i båt
func PutCornBoat() {
	state.CornBoat = true
	state.CornEast = false
	state.CornWest = false
}

//Setter HS i båt
func PutHsBoat() {
	state.HsBoat = true
	state.HsEast = false
	state.HsWest = false
}

//Setter Kylling i Land Vest
func PutChickenWest() {
	state.ChickenBoat = false
	state.ChickenEast = false
	state.ChickenWest = true
}

//Setter Rev i Land Vest
func PutFoxWest() {
	state.FoxBoat = false
	state.FoxEast = false
	state.FoxWest = true
}

//Setter Korn i Land Vest
func PutCornWest() {
	state.CornBoat = false
	state.CornEast = false
	state.CornWest = true
}

//Setter Hs i Land Vest
func PutHsWest() {
	state.HsBoat = false
	state.HsEast = false
	state.HsWest = true
}

//Setter Kylling i Land Øst
func PutChickenEast() {
	state.ChickenBoat = false
	state.ChickenEast = true
	state.ChickenWest = false
}

//Setter Rev i Land Øst
func PutFoxEast() {
	state.FoxBoat = false
	state.FoxEast = true
	state.FoxWest = false
}

//Setter Korn i Land Øst
func PutCornEast() {
	state.CornBoat = false
	state.CornEast = true
	state.CornWest = false
}

//Setter Hs i Land Øst
func PutHsEast() {
	state.HsBoat = false
	state.HsEast = true
	state.HsWest = false
}
