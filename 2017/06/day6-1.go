package main

import "fmt"

func main() {
	banks := []int{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
	//banks := []int{0, 2, 7, 0}
	history := make(map[string]bool)
	c := 0
	for {
		c++

		max, maxAt := 0, 0
		for i, v := range banks {
			if v > max {
				maxAt = i
				max = v
			}
		}

		fmt.Printf("Max %d at %d\n", max, maxAt)
		banks[maxAt] = 0
		for i := 0; i < max; i++ {
			banks[(maxAt+i+1)%len(banks)]++
		}

		rep := fmt.Sprintf("%v", banks)
		fmt.Printf("%d: %v\n", c, banks)
		if _, ok := history[rep]; ok == true {
			break
		}

		history[rep] = true
	}
	fmt.Printf("Count: %d\n", c)
}
