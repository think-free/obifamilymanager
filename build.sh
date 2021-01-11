#!/bin/bash

sudo docker build --rm -t thinkfree84/obifamilymanager . &&
sudo docker image prune -f --filter label=stage=intermediate
