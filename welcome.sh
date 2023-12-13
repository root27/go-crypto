#!/bin/bash

GREEN="\e[32m"
ENDCOLOR="\e[0m"


# Simulate a loading effect
chars="

 ################## Welcome to Coin App ##################
 

    Usage with arguments:
    "

greenChars=" 
        --help, -help: show this help message and exit
        --coin, -coin [ARG]: show coin info

            Exampe: -coin Bitcoin
        
        --all, -all: show first 500 coins
        --number, -number [ARG]: Number of coins to display. Use with -all flag 
        
            Example: -number 50 -all or -all -number 50"

for (( j=0; j<${#chars}; j++ )); do
    printf "${chars:$j:1}"
    sleep 0.005
done

for (( j=0; j<${#greenChars}; j++ )); do
    printf "${GREEN}${greenChars:$j:1}${ENDCOLOR}"
    sleep 0.005
done

printf "\n"