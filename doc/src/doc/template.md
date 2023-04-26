---
title: Template file
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
# Template file

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