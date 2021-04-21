#!/usr/bin/env bash

set -e
echo "" > coverage.out

echo -e "TESTS FOR: for \033[0;32mTDBackend\033[0m"
go-acc -o coverage.out ./... -- -timeout 30m