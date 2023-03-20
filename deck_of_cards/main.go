package main

//import "fmt"

func main() {

	//(1) newdeck := []string{"ace of spades", "two of diamonds"}
	// (1)newdeck = append(newdeck, "three of clubs") /// moved the functionality to deck.go
	// printDeck(newdeck)// ===>(2)===> printDeck(createNewDeck())

	// (1)func printDeck(d []string) { ===>(2)==> )func printDeck(d deck) // converted []string to deck in deck.go
	//(1)	fmt.Println(d)// (3) == copies whole function and pasted in deck.go
	//return (newdeck)
	// }

	newDeck := createNewDeck()

	//(3)//printDeck(deck) //(1)  // using it in object oriented way
	newDeck.print() //object oriented way
 

	// type conversion from []string to deck
	s := []string{"test"}
	deck(s).print()

}
