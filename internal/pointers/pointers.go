/* The exmaples used in this script have been obtains from the Digital Ocean
Course zxhttps://www.digitalocean.com/community/conceptual-articles/understanding-pointers-in-go
*/

package pointers

import "fmt"

func ModifyingPointers(creature string, new_creature string) (result string) {

	var pointer *string = &creature

	fmt.Printf("Creature = %v\n", creature)
	fmt.Printf("Pointer = %v\n", pointer)
	// Dereference the pointer variable
	fmt.Printf("*pointer value = %v\n", *pointer)
	// Modifying the value stored in the pointer variable's location
	*pointer = new_creature
	fmt.Printf("*pointer new value = %v\n", *pointer)

	return *pointer
}

// Function Pointer Receivers objects

type Creature struct {
	Species string
}

func FunctionPointerReceivers() {
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
