package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

func reactUnits(p []byte) []byte {
	var newP []byte

	for srcI := 0; srcI < len(p); {
		if len(newP) == 0 {
			newP = append(newP, p[srcI])
			srcI++
		}

		a, b := rune(newP[len(newP)-1]), rune(p[srcI])
		if (unicode.IsLower(a) && unicode.IsUpper(b) && a == unicode.ToLower(b)) ||
			(unicode.IsUpper(a) && unicode.IsLower(b) && a == unicode.ToUpper(b)) {
			srcI++
			newP = newP[:len(newP)-1]
		} else {
			newP = append(newP, p[srcI])
			srcI++
		}
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
		//fmt.Printf("Shrank from %d to %d\n", len(polymer), len(newPolymer))
		if len(newPolymer) == len(polymer) {
			break
		}
		polymer = newPolymer
	}

	fmt.Printf("Final length is: %d\n", len(polymer))
}
