# Crocc

Crocc is a simple static-site generator based on Markdown.
The main goal of Crocc is to offer simplicity, as opposed to other static-site
generation tools such as Hugo or Jekyll.
Indeed, you only need Markdown to write content in a productive manner, and
everything else is standard scripts (JavaScript, CSS, images, etc.).

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
$ crocc -out=dst -url="http://example.com" -sitemap src
```

The `dst` directory will contain the following files:

```
dst/
├── index.html
├── about.html
├── bar.png
├── contact.html
└── sitemap.xml
```

You can now upload the `dst` directory to your web server and you're done!

## Documentation

### Input/output directory

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

### Markdown document

The Markdown document must have a YAML header, also known as "front matter".
The YAML header is a set of key-value pairs separated by a colon.
The YAML header is followed by the Markdown document.

Front matter keys:
* `title`: The title of the document. Required.
* `description`: The description of the document. Required.
* `publication_time`: The date of the document. Required.
* `last_update_time`: The date of the last update of the document. Not required.
* `keywords`: The tags of the document, as a list of strings. Required.
* `author`: The author of the document. Default is `""`.
* `hide`: If set to `true`, the document will not be generated.
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
```

### Page template

To create HTML pages, Crocc uses a template file.
The template file, written in HTML using Go template syntax, must be located in
the `$INPUT/.crocc.html` file.

The template file is injected with a set of variables. A variable can be used
in the template file using the `{{ .VariableName }}` syntax.
Read the [Go template documentation](https: //golang.org/pkg/text/template) for
more information.

The following variables are available:
* `.Title`: The title of the document.
* `.Description`: The description of the document.
* `.PublicationTime`: The date of the document.
* `.LastUpdateTime`: The date of the last update of the document.
* `.Keywords`: The tags of the document, as a string separated by commas.
* `.Author`: The author of the document.
* `.Content`: The content of the document, as HTML.
* `.Site`: The URL of the site.
* `.Generator`: A string containing the name and version of the generator.
* `.Sitemap`: The URL of the sitemap.

Here is a sample template:

```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>{{ .Title }}</title>
	<meta name="description" content="{{ .Description }}">
	<meta name="keywords" content="{{ .Keywords }}">
	<meta name="author" content="{{ .Author }}">
	<meta name="generator" content="{{ .Generator }}">
	<meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
	<header>
		<h1>{{ .Title }}</h1>
	</header>
	<main>
		{{ .Content }}
	</main>
	<footer>
		<p>
			Last update: {{ .LastUpdateTime }}
		</p>
	</footer>
</body>
</html>
```

## License

The project is governed by a BSD-style license that can be found in the 
[LICENSE](LICENSE) file.