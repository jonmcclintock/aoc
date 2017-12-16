#!/bin/bash

BANKS=(14 0 15 12 11 11 3 5 1 6 8 4 9 1 8 4)
#BANKS=(0 2 7 0)
HISTORY=()
C=0
L=${#BANKS[@]}
FIRST_SEEN=0

while :
do
    ((C++))
    MAX=0
    MAX_AT=0
    for i in `seq 0 $L`; do
        v=${BANKS[$i]}
        if [[ $v -gt $MAX ]] ; then
            MAX_AT=$i
            MAX=$v
        fi
    done

    BANKS[$MAX_AT]=0
    ((POS=MAX_AT+i))
    for i in `seq 1 $MAX` ; do 
        ((POS=(POS+1) % L))
        ((BANKS[POS]=BANKS[POS]+1))
    done

    IFS=$'\n'
    REP="${BANKS[@]}"
    MATCH=`echo "${HISTORY[*]}" | grep -n "^$REP$"`
    if [[ $? -eq 0 ]]; then
        FIRST_SEEN=`echo $MATCH | sed 's/:.*//'`
        break
    fi

    HISTORY[${#HISTORY[@]}]="${BANKS[@]}"
done

echo Loop detected at: $C, first seen at $FIRST_SEEN.
