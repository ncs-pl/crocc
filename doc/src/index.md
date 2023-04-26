---
title: Crocc
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
publication_time: 2023-04-26T19:05:00Z
publication_time: 2023-04-27T00:08:00Z
author: Nicolas Paul <n@nc0.fr>
---
# Crocc

Crocc is a simple static-site generator based on Markdown.
The main goal of Crocc is to offer simplicity, as opposed to other static-site
generation tools such as Hugo or Jekyll.
Indeed, you only need Markdown to write content in a productive manner, and
everything else is standard scripts (JavaScript, CSS, images, etc.).

![A crocodile walking](/assets/crocodile.jpg)

## Installation

To install Crocc from sources, you need to have Go installed on your system.
Then, run the following command:

```bash
$ go install go.nc0.fr/crocc@latest
```

> Note: You can replace `latest` with a specific Git commit.

## Usage

See the [example](/example) for a simple example of a Crocc website.

A more concrete project is 
[this website itself](https://github.com/n1c00o/crocc/tree/master/doc), which 
is built with Crocc.

## Documentation

See the [documentation](/doc) for more information.

## License

The project is governed by a BSD-style license that can be found in the
[LICENSE](https://github.com/n1c00o/crocc/blob/master/LICENSE) file.

The banner image is a [photo](https://unsplash.com/photos/R3sgrDvXz3I) from 
[Unsplash](https://unsplash.com/) by 
[Thomas Couillard](https://unsplash.com/@thomascouillard).

Favicons are Noto emojis and are licensed under the 
[Apache License 2.0](https://github.com/googlefonts/noto-emoji/blob/main/LICENSE)
license.
