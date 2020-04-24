#!/usr/bin/env bash

function override_function() {
    local ORIG_FUNC=$(declare -f $1)
    local NEWNAME_FUNC="$2${ORIG_FUNC#$1}"
    eval "$NEWNAME_FUNC"
}


function test {
    echo "Echo value to terminal" >&2
}
function test_2 {
    echo "__________" >&2
}

override_function test_2 test
 test

#test

$*