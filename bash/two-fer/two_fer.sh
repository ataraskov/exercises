#!/usr/bin/env bash

do_two_fer() {
    echo "One for ${1:-you}, one for me."
}

do_two_fer "$@"