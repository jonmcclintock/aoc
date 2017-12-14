#!/bin/bash

IN=(14 0 15 12 11 11 3 5 1 6 8 4 9 1 8 4)
HIST=()
C=0
FIRST_SEEN=0

while :
do
    ((C=C+1))
    MAX=0
    AT=0
    for i in `seq 0 ${#IN[@]}`; do
        v=${IN[$i]}
        if [[ $v -gt $MAX ]] ; then
            AT=$i
            MAX=$v
        fi
    done

    IN[$AT]=0
    ((POS=AT+i))
    for i in `seq 1 $MAX` ; do 
        ((POS=(POS+1)%${#IN[@]}))
        ((IN[POS]=IN[$POS]+1))
    done

    IFS=$'\n'
    REP="${IN[@]}"
    MATCH=`echo "${HIST[*]}" | grep -n "^$REP$"`
    if [[ $? -eq 0 ]]; then
        FIRST_SEEN=`echo $MATCH | sed 's/:.*//'`
        break
    fi

    HIST[${#HIST[@]}]="${IN[@]}"
done

echo Loop detected at: $C, first seen at $FIRST_SEEN.
