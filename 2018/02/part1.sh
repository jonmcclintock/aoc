#!/bin/bash

function check_for_letter_count() {
    IN=$1
    TARGET=$2

    for LETTER in {a..z}; do
        OUT=`echo $IN | sed "s/$LETTER//g"`
        if (((${#IN} - ${#OUT}) == $TARGET)); then
            #echo "Found a $TARGET letter match for '$LETTER'"
            return 0
        fi
    done

    #echo "No match for count $TARGET found"
    return 1
}

COUNT_2=0
COUNT_3=0
for ID in `cat $1` ; do
    echo "Checking: $ID"
    if check_for_letter_count $ID 2; then
        ((COUNT_2++))
    fi
    if check_for_letter_count $ID 3; then
        ((COUNT_3++))
    fi
    echo "COUNT_2: $COUNT_2 COUNT_3: $COUNT_3"
done

echo "Count of 2-letter: $COUNT_2"
echo "Count of 3-letter: $COUNT_3"
echo "Product:" $((COUNT_2 * COUNT_3))
