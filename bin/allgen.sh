#!/bin/bash
set -eux

for tool in tools/*/*.go; do
    go run "$tool"
done

for script in bin/*gen.sh; do
    if [[ "$script" != "bin/allgen.sh" ]]; then
        bash "$script"
    fi
done
