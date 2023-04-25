// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crocc is a simple Markdown-based static site generator.
package main /* import "go.nc0.fr/crocc" */

import (
	"flag"
	"fmt"
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

// versionFormat returns a string containing the build information.
func versionFormat() string {
	return fmt.Sprintf("crocc version %s %s/%s %s date %s",
		version, runtime.GOOS, runtime.GOARCH, runtime.Compiler, date)
}

func main() {
	flag.Parse()

	log.SetFlags(0)

	if *printVersion {
		log.Print(versionFormat())
		return
	}

	inputdir := flag.Arg(0)
	if inputdir == "" {
		log.Fatalln("no input directory specified")
	}

	if *verbose {
		log.Printf(`Version: %s
Input directory: %s
Output directory: %s
Site URL: %s
Generate sitemap: %t
Generate hidden pages: %t`, versionFormat(), inputdir, *outputdir, *url,
			*sitemap, *generateHidden)
	}
}
