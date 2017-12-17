#!/bin/bash
ROUNDS=50000000; STEPS=301; CUR=0; RES=0
for i in `seq -f '%.0f' 1 $ROUNDS`; do ((CUR = (CUR + STEPS) % i)); if [[ $CUR == 0 ]] ; then RES=$i; fi; (( CUR++ )); done; echo After zero is $RES
