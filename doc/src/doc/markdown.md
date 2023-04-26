---
title: Markdown syntax
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
# Markdown document

The Markdown document must have a YAML header, also known as "front matter".
The YAML header is a set of key-value pairs separated by a colon.
The YAML header is followed by the Markdown document.

Front matter keys:

- `title`: The title of the document. Required.
- `description`: The description of the document. Required.
- `publication_time`: The date of the document. Required.
- `last_update_time`: The date of the last update of the document. Not required.
- `keywords`: The tags of the document, as a list of strings. Required.
- `author`: The author of the document. Default is `""`.
- `hide`: If set to `true`, the document will not be generated. 
  Default is `false`.

Example:

```md
---
title: Hello World
description: This is a simple example of a Markdown document.
publication_time: 2020-01-01T00:00:00Z
last_update_time: 2020-01-01T03:00:00Z
keywords: [example, hello, world]
author: John Doe
hide: true
---
# Hello World!

This is a simple example of a Markdown document.
```