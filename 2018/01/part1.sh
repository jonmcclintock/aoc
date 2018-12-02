#!/bin/bash

SUM=0
for VAL in `cat $1` ; do 
    (( SUM = SUM + VAL ));
    echo "Added $VAL, result is $SUM"
done

echo "Sum is $SUM"
