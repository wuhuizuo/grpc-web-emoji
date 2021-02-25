
#!/bin/bash

function verlte() {
    [  "$1" = "`echo -e "$1\n$2" | sort -V | head -n1`" ]
}

function verlt() {
    [ "$1" = "$2" ] && return 1 || verlte $1 $2
}

function red_print() {
    echo -ne "$(tput setaf 1)$1$(tput sgr0)"
}

function green_print() {
    echo -ne "$(tput setaf 2)$1$(tput sgr0)"
}

function yellow_print() {
    echo -ne "$(tput setaf 3)$1$(tput sgr0)"
}

function red_println() {
    echo -e "$(tput setaf 1)$1$(tput sgr0)"
}

function green_println() {
    echo -e "$(tput setaf 2)$1$(tput sgr0)"
}

function yellow_println() {
    echo -e "$(tput setaf 3)$1$(tput sgr0)"
}