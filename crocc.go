// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crocc is a simple Markdown-based static site generator.
package main /* import "go.nc0.fr/crocc" */

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

var (
	out            = flag.String("out", "dst", "output directory")
	url            = flag.String("url", "http://localhost", "site URL")
	generateHidden = flag.Bool("hidden", false, "generate hidden pages")
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

var (
	in           string
	htmlTemplate template.Template
)

func init() {
	flag.Usage = func() {
		fmt.Println(usage)
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

	// Check input directory
	in = flag.Arg(0)
	if in == "" {
		log.Fatalln("no input directory specified")
	}

	if _, err := os.Stat(in); os.IsNotExist(err) {
		log.Fatalf("input directory %q does not exist", in)
	}

	// Check output directory
	if _, err := os.Stat(*out); !os.IsNotExist(err) {
		log.Fatalf("output directory %q already exists", *out)
	} else {
		if err := os.MkdirAll(*out, 0755); err != nil {
			log.Fatalf("unable to create output directory %q: %v", *out, err)
		}
	}

	// Retrieve template file
	templatePath := filepath.Join(in, ".crocc.html")
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		log.Fatalf("template file %q does not exist", templatePath)
	}
	tp, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("error reading template file %q: %v", templatePath, err)
	}
	htmlTemplate = *template.Must(template.New("html-template").Parse(string(tp)))

	// Logic
	if err := filepath.WalkDir(in, Crocc); err != nil {
		log.Fatalf("unable to complete generation from %q: %v", in, err)
	}
}

// Crocc is the function that applies to every file in a directory.
// child corresponds to the path of a nested subdirectory, relative to the root.
// For example, if the root is "src" and the child is "foo/bar", the function
// will be applied to "src/foo/bar".
func Crocc(path string, d os.DirEntry, e error) error {
	if e != nil {
		return e
	}

	// ignore the input directory itself and hidden files
	if path == in || strings.HasPrefix(d.Name(), ".") {
		return nil
	}

	o := strings.Replace(path, in, *out, 1)

	// If the file is a directory, create it in the output directory
	if d.IsDir() {
		return TransformDirectory(o)
	}

	// Copy non-Markdown files into the output directory
	if filepath.Ext(path) != ".md" &&
		filepath.Ext(path) != ".markdown" &&
		filepath.Ext(path) != ".mdown" &&
		filepath.Ext(path) != ".Markdown" {
		return TransformNonMarkdownFile(path, o)
	}

	// If the file is a Markdown file, transform it into HTML
	o = strings.TrimSuffix(o, filepath.Ext(o)) + ".html"
	s := strings.TrimPrefix(o, *out)
	if err := TransformMarkdownFile(path, o, s); err != nil {
		return err
	}

	return nil
}
