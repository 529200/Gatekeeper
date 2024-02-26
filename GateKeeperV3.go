package main

import (
	"fmt"
	"time"
)

// V1 welkomst functie & V2 tijdsgebonden groet functie
func main() {
	fmt.Print(tijdcheck())
	fmt.Println("Welkom bij Fonteyn vakantieparken!")
	fmt.Print(kentekenControle())
}

// V2 Tijdsgebonden groet functie
func tijdcheck() string {
	var groet string
	tijd := time.Now()
	hour := tijd.Hour()

	if hour >= 7 && hour < 12 {
		groet = "Goedemorgen! "
	} else if hour >= 12 && hour < 18 {
		groet = "Goedemiddag! "
	} else if hour >= 18 && hour < 23 {
		groet = "Goedeavond! "
	} else {
		groet = "Sorry, de parkeerplaats is 's nachts gesloten! "
	}
	return groet
}

// V3 Kenteken controle functie
func kentekenControle() string {
	var paal string
	// hierbij wordt het kenteken van de auto gecheckt en daarbij kan hij zien of het kenteken in de registratie zit of niet.
	geautoriseerdekentekens := []string{"XYZ123", "ABC456", "DEF789", "ABCDEF"}
	var gescandekentekens string
	var geautoriseerd bool = false

	fmt.Print("Voer uw kenteken in: ")
	fmt.Scanln(&gescandekentekens)

	for _, geautoriseerdeplaat := range geautoriseerdekentekens {
		if gescandekentekens == geautoriseerdeplaat {
			geautoriseerd = true
			break
		}
	}

	if geautoriseerd {
		paal = "U heeft toegang tot het parkeerterrein"
	} else {
		paal = "U heeft helaas geen toegang tot het parkeer"
	}
	return paal

	// _ Wordt gebruikt zodat die geen index gaat maken van de kentekens die zijn gehardcode aangezien dit overbodig is voor ons gebruik.
}
