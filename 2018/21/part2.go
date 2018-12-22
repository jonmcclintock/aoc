package main

import "fmt"

func checkForCycle(history []int) bool {
	var cycleStart, cycleEnd int

search:
	for i := range history {
		for j := i + 1; j < len(history); j++ {
			if history[i] == history[j] {
				cycleStart = i
				cycleEnd = j - 1
				break search
			}
		}
	}

	if cycleStart == 0 {
		return false
	}

	minVal := history[cycleStart]
	for i := cycleStart; i <= cycleEnd; i++ {
		if history[i] < minVal {
			minVal = history[i]
		}
	}

	fmt.Printf("Found a cycle from %d to %d, smallest value is %d\n", cycleStart, cycleEnd, minVal)
	return true
}

func main() {
	previous := 0
	seen := make(map[int]bool, 0)
	regs := make([]int, 6)

	/*
	 5      seti 0 4 3              #three = 0
	 6      bori 3 65536 4          # $4 =three | 65536       // $4 = 0x10000
	 7      seti 1107552 3 3        #three = 1107552          //three = 0x10E660
	*/
	regs[3] = 0

line6:
	regs[4] = regs[3] | 65536
	regs[3] = 1107552

	/*
		 8      bani 4 255 5            # $5 = $4 & 255
		 9      addr 3 5 3              #three += $5
		10      bani 3 16777215 3       #three &= 16777215        // 0xFFFFFF
		11      muli 3 65899 3          #three *= 65899           //three = 72986569248
		12      bani 3 16777215 3       #three &= 16777215        //three = 5679648
	*/
line8:
	regs[5] = regs[4] & 255
	regs[3] += regs[5]
	regs[3] = ((regs[3] & 16777215) * 65899) & 16777215

	/*
		// if $4 <= 256 { goto 28 }
		13      gtir 256 4 5            # if 256 > $4 { $5 = 1 } else { $5 = 0 }
		14      addr 5 2 2              # Goto 14 + $5 + 1
		15      addi 2 1 2              # Goto 15 + 1 + 1
		16      seti 27 0 2             # Goto 27 + 1
	*/
	if 256 > regs[4] {
		regs[5] = 1

		/*
			// ifthree != $0 { goto 6 }
			28      eqrr 3 0 5              # ifthree == $0 { $5 = 1 } else { $5 = 0 }
			29      addr 5 2 2              # Goto 29 + $5 + 1
			30      seti 5 8 2              # Goto 5 + 1
		*/
		if regs[3] != regs[0] {
			if _, ok := seen[regs[3]]; ok {
				fmt.Printf("Found a repeat %d, previous was %d\n", regs[3], previous)
				return
			}
			seen[regs[3]] = true
			previous = regs[3]
			regs[5] = 0
			goto line6
		}
		return
	}
	regs[5] = 0

	/*
		// for $5 = 0; ($5+1)*256 <= $4; $5++ {
		17      seti 0 2 5              # $5 = 0
		18      addi 5 1 1              # $1 = $5 + 1
		19      muli 1 256 1            # $1 *= 256
		20      gtrr 1 4 1              # if $1 > $4 { $1 = 1 } else { $1 = 0 }
		21      addr 1 2 2              # Goto 21 + $1 + 1
		22      addi 2 1 2              # Goto 22 + 1 + 1
		23      seti 25 3 2             # Goto 25 + 1
		24      addi 5 1 5              # $5++
		25      seti 17 3 2             # Goto 17 + 1
		// }
	*/
	//for regs[5] = 0; (regs[5]+1)*256 <= regs[4]; regs[5]++ {
	//}
	regs[5] = regs[4] >> 8

	/*
		26      setr 5 3 4              # $4 = $5

		// Goto 8
		27      seti 7 4 2              # Goto 7 + 1
	*/
	regs[4] = regs[5]
	goto line8

}
