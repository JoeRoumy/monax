#!/usr/bin/env bash

dockerd-entrypoint.sh > /var/log/docker.log 2>&1 &
su monax -c "exec $@"
