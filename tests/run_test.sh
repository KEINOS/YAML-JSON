#!/usr/bin/env bash

path_dir_test=$(cd $(dirname $0); pwd)

cat $path_dir_test/sample.yaml | docker run --rm -i yaml-json -json && \
cat $path_dir_test/sample.json | docker run --rm -i yaml-json -yaml
