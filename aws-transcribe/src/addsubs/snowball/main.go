package main

import (
	"fmt"

	"github.com/kljensen/snowball"
)

func main() {

	strings := []string{"levitate", "graduation", "digestion"}
	for _, v := range strings {
		stemmed, err := snowball.Stem(v, "english", true)
		if err == nil {
			fmt.Println(stemmed) // Prints "accumul"
		}
	}

}
