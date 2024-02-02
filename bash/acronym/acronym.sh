#!/usr/bin/env bash
do_acronym() {
    # $@ string(s)
    for i in "$@"
    do
      if echo "$i" | grep -Eq -- '([[:space:]]|-)'; then
         do_acronym $(echo "$i"|sed -e 's/-/ /g')
         continue
      fi
      # ignore non alphanumeric input
      i="$(echo "$i" | sed -Er 's/[^[:alnum:]]//g')"
      # print first letter/number in uppwer case
      echo -n "$(echo "$i" | cut -c1 | tr '[:lower:]' '[:upper:]')"
    done
}

# remove special chars before calling acronym
do_acronym $(echo "$*"|sed -Er 's/[?*]//g')