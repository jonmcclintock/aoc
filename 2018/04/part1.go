package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Usage: %s input-file", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	guards := make(map[int][]int, 0)

	curGuard := -1
	curStart := -1

	guardNumberRe := regexp.MustCompile(`.*Guard #(\d+) .*`)
	minuteRe := regexp.MustCompile(`.* \d\d:(\d\d)] .*`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("Line: %s\n", line)

		if strings.Contains(line, "Guard") {
			s := guardNumberRe.ReplaceAllString(line, "$1")
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Couldn't parse the guard number: %s", s)
			}

			curGuard = n
			if _, ok := guards[curGuard]; !ok {
				guards[curGuard] = make([]int, 60)
			}

			if curStart != -1 {
				log.Fatal("Got a new guard when the last sleep hasn't ended")
			}
			curStart = -1

		} else if strings.Contains(line, "falls asleep") {
			if curGuard == -1 {
				log.Fatal("Got a start with no guard")
			}
			if curStart != -1 {
				log.Fatal("Got a start when the last sleep hasn't ended")
			}

			s := minuteRe.ReplaceAllString(line, "$1")
			start, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Couldn't parse the start minute: %s", s)
			}
			curStart = start
			//fmt.Printf("Guard %d started sleep at %d\n", curGuard, curStart)

		} else if strings.Contains(line, "wakes up") {
			if curGuard == -1 {
				log.Fatal("Got a wake with no guard")
			}
			if curStart == -1 {
				log.Fatal("Got a wake up when the guard isn't asleep")
			}
			s := minuteRe.ReplaceAllString(line, "$1")
			end, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("Couldn't parse the end minute: %s", s)
			}

			//fmt.Printf("Guard %d started sleep at %d, woke up at %d\n", curGuard, curStart, end)
			for i := curStart; i < end; i++ {
				guards[curGuard][i]++
			}
			curStart = -1

		} else {
			log.Fatalf("Unknown input line: %s\n", line)
		}
	}

	worstGuard := -1
	maxSleep := 0
	worstMinute := -1
	for id, minutes := range guards {
		sleepSum := 0
		maxMinute := 0
		maxMinuteCount := 0

		fmt.Printf("Guard %d: ", id)
		for m, v := range minutes {
			fmt.Printf("%c", byte('0')+byte(v))
			sleepSum += v
			if v > maxMinuteCount {
				maxMinuteCount = v
				maxMinute = m
			}
		}
		fmt.Printf("\n")
		fmt.Printf("Guard %d: %d minutes asleep, mostly at %d\n", id, sleepSum, maxMinute)
		if sleepSum > maxSleep {
			worstGuard = id
			maxSleep = sleepSum
			worstMinute = maxMinute
		}
	}

	fmt.Printf("Guard %d slept %d minutes total, mostly at minute %d (product is %d)\n", worstGuard, maxSleep, worstMinute, worstGuard*worstMinute)
}
