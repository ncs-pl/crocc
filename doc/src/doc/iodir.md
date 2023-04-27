---
title: Input/output directory
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
publication_time: 2023-04-26T19:00:00Z
publication_time: 2023-04-26T23:57:43Z
author: Nicolas Paul <n@nc0.fr>
---
# Input/output directory

Crocc requires two directories to work: the input directory and the output 
directory.

## Input directory

The input directory is a folder containing all the files used to build the
site.
At its root, it should contain a `.crocc.html` file, which is the 
[template file](/doc/template).

> Only the top-level template is used. Nested templates are not supported.

Each Markdown file in this directory will be transformed to an HTML file.
Other files will be copied as-is to the output directory.
Crocc does not process hidden files – files whose name starts with a dot (`.`).

For illustration, here is a sample input directory:

![Sample input directory](/assets/doc_input_directory.png)

## Output directory

The output directory is a folder containing the generated site.
It is created by Crocc if it does not exist.

> Crocc will not delete the output directory if it already exists.
> Instead, it will exit.

The output directory is defined at compile time by using the `-out` 
[flag](/doc/cli#flags). If not specified, it defaults to `./dst`.

Here is the output directory corresponding to the previous input directory:

![Sample output directory](/assets/doc_output_directory.png)

As you can see, the output directory contains the same files as the input
directory, except for the template file and the Markdown files, which have been
converted to HTML documents.
