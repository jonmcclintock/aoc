#!/bin/bash

startA=699
startB=124
rounds=5000000
seedA=16807
seedB=48271
factor=2147483647
criteriaA=4
criteriaB=8

gen() {
	r=$3
	while :; do
		(( r=($1*r) % factor ))
		if (( (r % $2) == 0 )); then return $r; fi
	done
}

matches=0
a=startA
b=startB
for i in `seq 0 $rounds`; do
	generate $seedA $criteriaA $a
	a=$?
	generate $seedB $criteriaB $b
	b=$?

	if (( (a & 0xffff) == (b & 0xffff) )); then
		(( matches=matches+1 ))
	fi
done

echo Matches: $matches
