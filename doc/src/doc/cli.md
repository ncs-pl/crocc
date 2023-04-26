---
title: CLI
description: Crocc is a simple and fast static-site generator based on Markdown. 
keywords:
  - crocc
  - markdown
  - html
  - go
  - golang
  - static
  - site
  - generator
  - ssg
  - website
  - simple
publication_time: 2023-04-27T00:07:00Z
author: Nicolas Paul <n@nc0.fr>
---
# CLI

Crocc is a command-line tool.
It accepts various flags to customize its behavior.

## Usage

```bash
$ crocc [flags] [input directory]
```

## Input directory

The `[input directory]` argument is the path to the 
[input directory](/doc/iodir#input-directory) containing the files used to
build the site.

## Flags

Flags are used to customize the behavior of Crocc.

### `-out`

The `-out` flag is used to specify the 
[output directory](/doc/iodir#output-directory).

If not specified, it defaults to `./dst`.

### `-version`

The `-version` flag is used to print the current version of Crocc.
If specified, Crocc will print its version and exit.
The version string is formatted as:

```go
fmt.Sprintf("crocc version %s %s/%s %s date %s",
		version, runtime.GOOS, runtime.GOARCH, runtime.Compiler, date)
```

See the `versionFormat()` function in the 
[crocc.go](https://github.com/n1c00o/crocc/blob/master/crocc.go#L55-L59) file 
for the exact implementation.

### `-help`

The `-help` flag is used to print the help message and exit.

### `-hidden`

The `-hidden` flag is used to build all Markdown files, even those with
`hide: true` in their front matter.

See the [front matter](/doc/markdown#hidden-documents) documentation for more 
information.

### `-url`

The `-url` flag is used to specify the hostname of the site, e.g. 
`-url="https://example.com"`.
It can be used inside templates through the `{{ .Site }}` variable.

