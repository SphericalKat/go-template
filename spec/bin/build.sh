#!/bin/bash
# TODO: consider rewrite

# Params
set -eu
script_dir=$(cd $(dirname $0); pwd)


# Execution
cd ${script_dir}/../

for file in `find . -maxdepth 2 -name index.yaml -type f`; do
    site=$(dirname $file | sed -e 's/\.\///g')
    build_dir="../build/${site}"

    pushd $site
        mkdir -p ${build_dir}/
        swagger-cli bundle index.yaml -o ${build_dir}/index.yaml -t yaml
        # redocly build-docs ${build_dir}/index.yaml -o ${build_dir}/index.html
    popd
done