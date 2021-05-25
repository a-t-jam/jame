#!/usr/bin/env bash
#
# Generates `statik` directory

# Run `go get github.com/rakyll/statik` to install `statik.
# Make sure `$GOPATH/bin` is in your path (so that `statik` is available).

statik -src="$(pwd)/assets" -include=*.jpg,*.png,*.wav,*.mp3

