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

func stripUnits(p []byte, c rune) []byte {
	var newP []byte
	r := unicode.ToLower(c)

	for i := 0; i < len(p); i++ {
		if r != unicode.ToLower(rune(p[i])) {
			newP = append(newP, p[i])
		}
	}

	return newP
}

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s input-file", os.Args[0])
	}

	basePolymer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lengths := make(map[rune]int, 0)

	shortest := ' '
	lengths[shortest] = len(basePolymer)

	for c := 'a'; c <= 'z'; c++ {
		polymer := stripUnits(basePolymer, c)

		for {
			newPolymer := reactUnits(polymer)
			if len(newPolymer) == len(polymer) {
				break
			}
			polymer = newPolymer
		}

		fmt.Printf("Final length for '%c' is: %d\n", c, len(polymer))
		lengths[c] = len(polymer)

		if len(polymer) < lengths[shortest] {
			shortest = c
		}
	}

	fmt.Printf("Shortest is from removing '%c', resulting in %d units\n", shortest, lengths[shortest])

}
