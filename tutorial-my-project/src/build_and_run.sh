#!/bin/bash
# chmod +x build_and_run.sh

dir=$(dirname "$0")
cd "$dir" || exit

go mod tidy

go build -o rssagg && ./rssagg
