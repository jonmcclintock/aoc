#!/bin/bash

startA=699         # Starting value for generator A
startB=124         # Starting value for generator B
rounds=5000000     # Number of rounds to perform
factorA=16807      # Generator A's factor
factorB=48271      # Generator B's factor
limit=2147483647   # Modulo for multiplication
criteriaA=4        # Generator A results must divide evenly by this
criteriaB=8        # Generator B results must divide evenly by this

gen() {
	r=$3
	while :; do
		(( r=($1*r) % limit ))
		if (( (r % $2) == 0 )); then return; fi
	done
}

a=$startA
b=$startB
matches=0
for i in `seq 0 $rounds`; do
	gen $factorA $criteriaA $a; a=$r
	gen $factorB $criteriaB $b; b=$r

	if (( (a & 0xffff) == (b & 0xffff) )); then
		(( matches++ ))
	fi
done

echo Matches: $matches
