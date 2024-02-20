/* The exmaples used in this script have been obtains from the Digital Ocean
Course zxhttps://www.digitalocean.com/community/conceptual-articles/understanding-pointers-in-go
*/

package pointers

import "fmt"

func UsingPointers() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Printf("Creature = %v\n", creature)
	fmt.Printf("Pointer = %v\n", pointer)
	// Dereference the pointer variable
	fmt.Printf("*pointer value = %v\n", *pointer)
	// Modifying the value stored in the pointer variable's location
	*pointer = "jellyfish"
	fmt.Printf("*pointer new value = %v\n", *pointer)
}
