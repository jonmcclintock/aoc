package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

func reactUnits(p []byte) []byte {
	var i int
	var newP []byte

	for i := 0; i < len(p)-1; i++ {
		a, b := rune(p[i]), rune(p[i+1])
		if (unicode.IsLower(a) && unicode.IsUpper(b) && a == unicode.ToLower(b)) ||
			(unicode.IsUpper(a) && unicode.IsLower(b) && a == unicode.ToUpper(b)) {
			i++
			continue
		}
		newP = append(newP, byte(a))
	}

	if i < len(p) {
		newP = append(newP, p[i])
	}

	return newP
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s input-file", os.Args[0])
	}

	polymer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for {
		newPolymer := reactUnits(polymer)
		fmt.Printf("Shrank from %d to %d\n", len(polymer), len(newPolymer))
		if len(newPolymer) == len(polymer) {
			break
		}
		polymer = newPolymer
	}

	fmt.Printf("Final length is: %d\n", len(polymer))
}
