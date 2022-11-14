#!/usr/bin/env bash

do_help() {
    echo "Usage: leap.sh <year>"
}
do_leap() {
    # not single argument
    if [ $# -ne 1 ]; then
      do_help
      return 1
    fi

    # non numeric argument
    if ! echo "$1"|grep -qE '^[[:digit:]]+$'; then
      do_help
      return 1
    fi

    date -d "$1-02-29" >/dev/null 2>&1 && echo "true" || echo "false"
}

do_leap "$@"