//go:build ignore

package main

import (
	"fmt"

	"github.com/RaphaelPour/silo"
)

func main() {
	// create a new file-based silo
	store := silo.NewCache(silo.NewJson(silo.NewFile("data.store")))

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
