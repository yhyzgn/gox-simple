#!/usr/bin/env bash
PACKAGE="gox-simple/internal"
VERSION=#version#
BUILD=$(TZ=UTC-8 date "+%Y-%m-%d_%H:%M:%S")

CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static' -X $PACKAGE/built.Version=$VERSION -X $PACKAGE/built.BuildDate=$BUILD"
