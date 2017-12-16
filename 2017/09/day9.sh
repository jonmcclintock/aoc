#!/bin/bash

IN=`cat day9-in.txt`

T=0; D=0; S="N"; G=0; for (( i=0; i < ${#IN}; i++ )); do case $S in "N") case ${IN:$i:1} in '{') ((D++)) ;; '}') ((T+=D)); ((D--)) ;; '<') S="G" ;; esac ;; "G") case ${IN:$i:1} in '!') S="E" ;; '>') S="N" ;; *) ((G++)) ;; esac ;; "E") S="G" ;; esac; done; echo Total: $T Garbage: $G
