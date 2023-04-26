#!/usr/bin/env bash
# Copyright (c) 2023 Nicolas Paul All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


set -e
rm -rf testdata/dst
./crocc -out=testdata/dst testdata/src
