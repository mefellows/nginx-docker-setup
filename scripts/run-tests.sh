#!/bin/sh

cd $(dirname $0)/../test
docker-compose stop && docker-compose  rm -f && docker-compose build && docker-compose  up --force-recreate
