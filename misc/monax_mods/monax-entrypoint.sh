#!/usr/bin/env bash

export MONAX_PULL_APPROVE=true
su monax -c "monax init --yes --pull-images=false"
su monax -c "exec $@"
