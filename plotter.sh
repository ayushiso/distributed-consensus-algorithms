#!/usr/bin/env bash

# Takes .tsv matrix files, runs proportions.py and saves the results to working directory 

for i in * 
do
    if [[ $i == *".tsv" ]] ; then
        echo "Starting '$i'"
        python proportions.py $i
    fi
done