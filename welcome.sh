#!/bin/bash

GREEN="\e[32m"
ENDCOLOR="\e[0m"


# Simulate a loading effect
chars="

 ################## Welcome to Coin App ##################
 

    Usage with arguments:
    "

greenChars=" 
        --help: show this help message and exit
        --coin [ARG]: show coin info
        --all: show first 10 coins"


for (( j=0; j<${#chars}; j++ )); do
    printf "${chars:$j:1}"
    sleep 0.005
done

for (( j=0; j<${#greenChars}; j++ )); do
    printf "${GREEN}${greenChars:$j:1}${ENDCOLOR}"
    sleep 0.005
done

printf "\n"