#!/usr/bin/env bash

do_reverse() {
    local word="$1"
    while [ -n "$word" ]
    do
        echo -n "${word: -1}"
        word="${word::-1}"
    done
}

do_reverse "$1"