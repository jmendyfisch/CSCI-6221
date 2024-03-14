#!/bin/bash

dcoker login --username swarupgt
docker tag asp_server:latest swarupgt/server-docker:latest
docker push swarupgt/server-docker:latest