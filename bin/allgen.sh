#!/bin/bash
set -eux

for script in bin/*gen.sh; do
    if [[ "$script" != "bin/allgen.sh" ]]; then
        bash "$script"
    fi
done
