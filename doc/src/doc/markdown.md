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
last_update_time: 2023-04-27T00:04:00Z
author: Nicolas Paul <n@nc0.fr>
---
# Markdown document

A Markdown document is a file whose extension is either `.md`, `.markdown`,
`.mdown` or `.Markdown`.

Markdown documents are used to generate HTML documents.
They should be placed in the [input directory](/doc/iodir#input-directory).
At the top of the Markdown document, there must be a YAML header.
See the [front matter](#front-matter) section for more information.

## Markdown syntax

Crocc uses [`github.com/gomarkdown/markdown`](https://github.com/gomarkdown/markdown)
(*GoMarkdown*) to parse Markdown documents.

GoMarkdown is a Markdown parser written in Go that supports various Markdown
extensions.
By default, GoMarkdown supports the CommonMark specification.

Crocc also uses the following extensions of GoMarkdown:

```go
extensions := parser.Tables | parser.FencedCode |
		parser.Autolink | parser.Strikethrough | parser.SpaceHeadings |
		parser.HeadingIDs | parser.BackslashLineBreak |
		parser.AutoHeadingIDs | parser.Footnotes | parser.SuperSubscript |
		parser.NoIntraEmphasis
```

- `parser.Tables` enables the parsing of tables similar to the ones in
  [GitHub Flavored Markdown](https://github.github.com/gfm/#tables-extension-).
- `parser.FencedCode` enables the parsing of fenced code blocks and generating 
  HTML `<code>` tags to help with syntax highlighting.
- `parser.Autolink` enables the parsing of URLs that are not explicitly marked.
- `parser.Strikethrough` enables the parsing of strike-through text using two 
  tildes (`~~test~~`).
- `parser.SpaceHeadings` makes the parser more strict about prefix heading 
  rules.
- `parser.HeadingIDs` enables the parsing of custom heading IDs using `{#id}`.
  For instance, the following will generate a heading with the ID
  `hello` instead of the default `hello-world`:

  ```md
  {#hello}
  # Hello World 
  ```
- `parser.BackslashLineBreak` enables the parsing of trailing backslashes as
  line breaks, as present in the original `Markdown.pl` implementation.
- `parser.AutoHeadingIDs` enables the generation of heading IDs from the text.
  For instance, the following will generate a heading with the ID
  `hello-world`:

  ```md
  # Hello World 
  ```
- `parser.Footnotes` enables the parsing of footnotes using the following
  syntax:

  ```md
  Here is a footnote reference,[^1] and another.[^longnote]

  [^1]: Here is the footnote.
  [^longnote]: Here's one with multiple blocks.

      Subsequent paragraphs are indented to show that they
  belong to the previous footnote.
  ```
- `parser.SuperSubscript` enables the parsing of super- and subscript text
  using `^` and `~` respectively.
  For illustration, the following will generate `H<sub>2</sub>O` and
  `2<sup>10</sup>`:

  ```md
  H~2~O
  2^10
  ```
- `parser.NoIntraEmphasis` disables the parsing of emphasis inside of words,
  as in `foo_bar_baz`.

## Front matter

The Markdown document must have a YAML header, also known as "front matter".
The YAML header is a set of key-value pairs separated by a colon, placed inside
two lines of three dashes (`---`).

Front matter keys:

| Key                |      Type       | Description                                                                             |
| ------------------ | :-------------: | --------------------------------------------------------------------------------------- |
| `title`            |     string      | The title of the document. Required.                                                    |
| `description`      |     string      | The description of the document. Required.                                              |
| `publication_time` |     string      | The date of the document. Required.                                                     |
| `last_update_time` |     string      | The date of the last update of the document. Not required.                              |
| `keywords`         | list of strings | The tags of the document, as a list of strings. Required.                               |
| `author`           |     string      | The author of the document. Default is `""`.                                            |
| `hide`             |     boolean     | If set to `true`, the document will be [hidden](#hidden-documents). Default is `false`. |

For instance:

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

Front matter keys are used to populate the [template file](/doc/template),
thus allowing the user to customize the generated HTML document.

## Hidden documents

A document can be hidden by setting the `hide` key to `true` in the front 
matter.
Hidden documents are not included in the generated site, unless the
[`-hidden` flag](/doc/cli#-hidden) is set.

Hidden documents are useful for drafts (documents that are not ready to
be published yet).

