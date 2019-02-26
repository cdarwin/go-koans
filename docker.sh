#!/usr/bin/env bash
docker run --rm -ti -v "$PWD":/usr/src/koans -w /usr/src/koans golang:1.6.0-alpine go test
