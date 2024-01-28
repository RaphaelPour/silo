package main

import (
	"fmt"

	"github.com/RaphaelPour/silo"
)

func main() {
	// create a new direct silo
	store := silo.NewDirect()

	// set a key
	err := store.Set("favorite-color", "purple")
	if err != nil {
		fmt.Println(err)
		return
	}

	// get key and print value
	rawValue, err := store.Get("favorite-color")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, ok := rawValue.(string)
	if !ok {
		fmt.Println("value is not a string")
		return
	}

	fmt.Printf("favorite-color: %s\n", value)

	// delete key
	err = store.Delete("favorite-color")
	if err != nil {
		fmt.Println(err)
		return
	}
}
