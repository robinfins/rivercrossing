package main

import (
	"bufio"
	"fmt"
	"os"
	"rivercrossing/event"
	"rivercrossing/state"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Print()

		//Status kommando
		if "status" == input.Text() {
			fmt.Println("Båt:       " + state.StateBoat())
			fmt.Println("Land Vest: " + state.StateLandV())
			fmt.Println("Land Øst:  " + state.StateLandE())
		}

		//Kylling kommandoer
		if "PutChickenBoat" == input.Text() {
			event.PutChickenBoat()
			fmt.Println("Kyllingen er i båten")
		}

		if "PutChickenWest" == input.Text() {
			event.PutChickenWest()
			fmt.Println("Kyllingen er på land vest")
		}

		if "PutChickenEast" == input.Text() {
			event.PutChickenEast()
			fmt.Println("Kyllingen er på land øst")
		}

		//Rev kommandoer
		if "PutFoxBoat" == input.Text() {
			event.PutFoxBoat()
			fmt.Println("Reven er i båten")
		}

		if "PutFoxWest" == input.Text() {
			event.PutFoxWest()
			fmt.Println("Reven er på land vest")
		}

		if "PutFoxEast" == input.Text() {
			event.PutFoxEast()
			fmt.Println("Reven er på land øst")
		}

		//Korn kommandoer
		if "PutCornBoat" == input.Text() {
			event.PutCornBoat()
			fmt.Println("Kornet er i båten")
		}

		if "PutCornWest" == input.Text() {
			event.PutCornWest()
			fmt.Println("Kornet er på land vest")
		}

		if "PutCornEast" == input.Text() {
			event.PutCornEast()
			fmt.Println("Kornet er på land øst")
		}

		//HS kommandoer
		if "PutHsBoat" == input.Text() {
			event.PutHsBoat()
			fmt.Println("HS er i båten")
		}

		if "PutHsWest" == input.Text() {
			event.PutHsWest()
			fmt.Println("HS er på land vest")
		}

		if "PutHsEast" == input.Text() {
			event.PutHsEast()
			fmt.Println("HS er på land øst")
		}
	}
}
