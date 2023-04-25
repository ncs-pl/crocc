# Copyright (c) 2023 Nicolas Paul All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

BINARY := crocc

VERSION := $(shell git describe --tags --always --dirty)
DATE := $(shell date -u '+%FT%T%z')

LDFLAGS := -ldflags "-w -s -X main.version=$(VERSION) -X main.date=$(DATE)"

$(BINARY): *.go
	go build $(LDFLAGS) -o $(BINARY)

clean:
	rm -f $(BINARY)

.PHONY: clean
