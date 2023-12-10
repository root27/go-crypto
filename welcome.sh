#!/bin/bash


# Simulate a loading effect
chars="################## Welcome to Coin App ##################
 ##########################################################
    Usage with arguments:
    -h, --help: show this help message and exit
    -c, --coin: coin name
    -a, --all: all coins
    "


for (( j=0; j<${#chars}; j++ )); do
    printf "${chars:$j:1}"
    sleep 0.005
done

printf "\n"