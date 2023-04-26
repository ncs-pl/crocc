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

Let's say you have a directory containing the following files:

```
src/
├── .crocc.html
├── index.md
├── about.md
├── bar.png
└── contact.md
```

The `.crocc.html` file is the template used to generate the HTML pages.
The `index.md`, `about.md` and `contact.md` files are Markdown documents.
The `bar.png` file is a static file.

To generate the HTML files, run the following command:

```bash
$ crocc -out=dst -url="http://example.com" src
```

The `dst` directory will contain the following files:

```
dst/
├── index.html
├── about.html
├── bar.png
└── contact.html
```

You can now upload the `dst` directory to your web server and you're done!

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
