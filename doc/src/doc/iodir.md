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
author: Nicolas Paul <n@nc0.fr>
---
# Input/output directory

The input directory is the directory containing the various files used to
build the site.
The input directory must contain a `.crocc.html` file, which is the
template used to generate the HTML pages.

> Only the top-level template is used. Nested templates are not supported.

The output directory is the directory where the generated HTML files will be
written.
The output directory must not exist before running Crocc.

Crocc will copy all the files in the input directory to the output directory,
except the `.crocc.html` file.
During the copy, Crocc will transform Markdown documents to HTML files.