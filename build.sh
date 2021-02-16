#!/bin/bash

touch .lastbuild
docker build --rm -t thinkfree84/obifamilymanager . &&
docker image prune -f --filter label=stage=intermediate

