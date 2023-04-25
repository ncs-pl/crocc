// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crocc is a simple Markdown-based static site generator.
package main /* import "go.nc0.fr/crocc" */

import (
	"flag"
	"log"
	"runtime"
)

var (
	outputdir      = flag.String("out", "dst", "output directory")
	url            = flag.String("url", "http://localhost", "site URL")
	sitemap        = flag.Bool("sitemap", false, "generate sitemap.xml")
	generateHidden = flag.Bool("hidden", false, "generate hidden pages")
	verbose        = flag.Bool("v", false, "verbose output")
	printVersion   = flag.Bool("version", false, "print version and exit")
)

const usage string = `crocc is a simple Markdown-based static site generator.
It is designed to be simple and fast, and to be used with a static web server.

Author: Nicolas Paul <n@nc0.fr> (https://nc0.fr)
License: BSD 3-Clause
Repository: https://github.com/n1c00o/crocc

Usage:
	crocc [options] [input directory]
	
Options:`

// Set at compilation time
var (
	version string
	date    string
)

func init() {
	flag.Usage = func() {
		log.Println(usage)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	log.SetFlags(0)

	if *printVersion {
		log.Printf("crocc version %s %s/%s %s date %s",
			version, runtime.GOOS, runtime.GOARCH, runtime.Compiler, date)
		return
	}
}
