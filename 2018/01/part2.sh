#!/bin/bash

CUR_FREQ=0
declare -a SEEN_FREQS

while true ; do
    for OFFSET in `cat $1` ; do
        (( CUR_FREQ = CUR_FREQ + OFFSET ));
        echo "Added $OFFSET, result is $CUR_FREQ"

        if echo ${SEEN_FREQS[@]} | grep "\b$CUR_FREQ\b" >> /dev/null ; then 
            echo "Found frequency $FREQ again, stopping."
            break 2
        fi

        SEEN_FREQS+=($CUR_FREQ)
    done
done

echo "Final frequency is: $CUR_FREQ"
