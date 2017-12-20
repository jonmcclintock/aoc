package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day20-in.txt"

type particle struct {
	pos           []int
	vel           []int
	acc           []int
	distanceDelta int
	dead          bool
}

func parseTuple(in string) []int {
	s := strings.Split(in[3:len(in)-1], ",")
	tuple := make([]int, 3)
	tuple[0], _ = strconv.Atoi(s[0])
	tuple[1], _ = strconv.Atoi(s[1])
	tuple[2], _ = strconv.Atoi(s[2])
	return tuple
}

func addTuple(a, b []int) []int {
	return []int{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func distanceFromOrigin(p []int) int {
	return int(math.Abs(float64(p[0])) +
		math.Abs(float64(p[1])) +
		math.Abs(float64(p[2])))
}

func loadParticles() []particle {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	particles := []particle{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ", ")
		p := particle{
			pos: parseTuple(s[0]),
			vel: parseTuple(s[1]),
			acc: parseTuple(s[2]),
		}
		particles = append(particles, p)
	}

	return particles
}

func dumpParticles(particles []particle) {
	for i, p := range particles {
		fmt.Printf(" %d: p=<%d,%d,%d>, v=<%d, %d, %d>, a=<%d,%d,%d> delta: %d, distance: %d, acc: %d, dead: %t\n", i,
			p.pos[0], p.pos[1], p.pos[2],
			p.vel[0], p.vel[1], p.vel[2],
			p.acc[0], p.acc[1], p.acc[2],
			p.distanceDelta,
			distanceFromOrigin(p.pos), distanceFromOrigin(p.acc),
			p.dead)
	}
}

func tick(particles []particle) bool {
	gettingCloser := false
	for i, p := range particles {
		if p.dead {
			continue
		}
		dStart := distanceFromOrigin(p.pos)
		p.vel = addTuple(p.vel, p.acc)
		p.pos = addTuple(p.pos, p.vel)
		dEnd := distanceFromOrigin(p.pos)
		delta := dEnd - dStart
		if delta <= p.distanceDelta {
			gettingCloser = true
		}
		p.distanceDelta = delta
		particles[i] = p
	}

	for i, p1 := range particles {
		if p1.dead {
			continue
		}
		for j, p2 := range particles {
			if i == j {
				continue
			}
			if p2.dead {
				continue
			}
			if p1.pos[0] == p2.pos[0] &&
				p1.pos[1] == p2.pos[1] &&
				p1.pos[2] == p2.pos[2] {
				fmt.Printf("%d (%v) and %d (%v) match, killing them\n", i, p1.pos, j, p2.pos)
				p1.dead = true
				p2.dead = true
				particles[i] = p1
				particles[j] = p2
			}
		}
	}

	return gettingCloser
}

func printClosest(particles []particle) {
	min := distanceFromOrigin(particles[0].pos)
	minAt := 0
	for i := range particles {
		d := distanceFromOrigin(particles[i].pos)
		if d < min {
			min = d
			minAt = i
		}
	}
	fmt.Printf("Closest is %d, with distance of %d\n", minAt, min)
}

func printMinAccel(particles []particle) {
	min := distanceFromOrigin(particles[0].acc)
	minAcc := []int{}
	for i := range particles {
		d := distanceFromOrigin(particles[i].acc)
		if d == min {
			minAcc = append(minAcc, i)
		}
		if d < min {
			min = d
			minAcc = []int{i}
		}
	}
	fmt.Printf("Found %d particles with min acc of %d\n", len(minAcc), min)

	min = distanceFromOrigin(particles[minAcc[0]].pos)
	minAt := minAcc[0]
	for i := range minAcc {
		d := distanceFromOrigin(particles[minAcc[i]].pos)
		if d < min {
			min = d
			minAt = minAcc[i]
		}
	}

	fmt.Printf("%d has smallest total distance %d at smallest acceleration\n", minAt, min)
}

func printSurvivorCount(particles []particle) {
	c := 0
	for _, p := range particles {
		if !p.dead {
			c++
		}
	}
	fmt.Printf("Survivors: %d\n", c)
}

func main() {
	particles := loadParticles()
	fmt.Printf("Starting positions:\n")
	dumpParticles(particles)
	printMinAccel(particles)

	for i := 1; ; i++ {
		fmt.Printf("Round %d:\n", i)
		if !tick(particles) {
			fmt.Printf("Not getting closer any more, stopping.\n")
			break
		}
		dumpParticles(particles)
	}
	fmt.Printf("Final state:\n")
	dumpParticles(particles)

	printClosest(particles)
	printMinAccel(particles)
	printSurvivorCount(particles)
}
