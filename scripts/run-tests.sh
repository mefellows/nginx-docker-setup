#!/bin/bash

cd $(dirname $0)/../test
BINARY="docker-compose"
ARG="-f "

function run() {
  RUN_COMMAND="${BINARY} ${ARG} ${1}.yml"
  ${RUN_COMMAND} stop && ${RUN_COMMAND}  rm -f && ${RUN_COMMAND} build && ${RUN_COMMAND}  up --force-recreate
}


if [ -z "$1" ]; then
  echo "Running entire test suite"
  run "perf"
  run "integration"
  run "muxy"
else
  echo "Running $1 test suite"
  run $1
fi
