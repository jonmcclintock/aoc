#!/bin/bash

BANKS=(14 0 15 12 11 11 3 5 1 6 8 4 9 1 8 4)
#BANKS=(0 2 7 0)
HISTORY=()
C=0
FIRST_SEEN=0

while :
do
    C=`expr $C + 1`
    MAX=0
    MAX_AT=0
    for i in `seq 0 ${#BANKS[@]}`; do
        if [[ ${BANKS[$i]} -gt $MAX ]] ; then
            MAX_AT=$i
            MAX=${BANKS[$i]}
        fi
    done

    echo Max $MAX at $MAX_AT

    BANKS[$MAX_AT]=0
    POS=`expr $MAX_AT + $i`
    for i in `seq 1 $MAX` ; do 
        POS=`expr \( $POS + 1 \) % ${#IN[@]}`
        BANKS[$POS]=`expr ${BANKS[$POS]} + 1`
    done

    echo $C: ${BANKS[@]}
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
