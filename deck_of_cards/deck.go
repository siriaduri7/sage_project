package main

import "fmt"

type deck []string // instead of []string we can use deck in all places so now replacing []string with deck (acts like laias)

func createNewDeck() deck { //(before) func createNewDeck() []string
	newdeck := deck{"ace of spades", "two of diamonds"}
	//newdeck = append(newdeck, "three of clubs")
	fmt.Println(newdeck)
	return newdeck

}

func (d deck) print() {
	//fmt.Println(d)
	for i, card := range d { // i is index, card is value , range is iterate in d, here d is slice
		fmt.Println(i, card)

	}
}
