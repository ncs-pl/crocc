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
last_update_time: 2023-04-27T00:07:00Z
author: Nicolas Paul <n@nc0.fr>
---
# Template file

To create dynamic HTML pages, Crocc uses a template file.
The template file, written in HTML using Go template syntax, must be located in
the `.crocc.html` file at the root of the 
[input directory](/doc/iodir#input-directory).

The template file is injected with a set of variables. 
A variable can be used in the template file using the `{{ .VariableName }}`
syntax.
Read the [Go template documentation](https://golang.org/pkg/text/template) for
more information.

The following variables are available:
- `.Title`: The title of the document.
- `.Description`: The description of the document.
- `.PublicationTime`: The date of the document.
- `.LastUpdateTime`: The date of the last update of the document.
- `.Keywords`: The tags of the document, as a string separated by commas.
- `.Author`: The author of the document.
- `.Content`: The content of the document, as HTML.
- `.Site`: The URL of the site.
- `.Generator`: A string containing the name and version of the generator.

> **Note:** The `.Content` variable is already HTML, so it does not need to be
> escaped.

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

Using templates, you can create a navigation bar, a footer, or any other
recurring element on your site.

For illustration purposes, here is the template used for this documentation:

```html
<!DOCTYPE html>
<!--
 Copyright (c) 2023 Nicolas Paul All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
-->
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>{{ .Title }}</title>
		<meta name="author" content="{{ .Author }}" />
		<meta name="description" content="{{ .Description }}" />
		<meta name="keywords" content="{{ .Keywords }}" />
		<meta name="generator" content="{{ .Generator }}">
		<meta name="color-scheme" content="light dark" />
		<meta name="theme-color" content="#155eca">
		<meta name="robots" content="all" />
		<link rel="sitemap" href="/sitemap.xml" />
		<meta name="twitter:card" content="summary_large_image" />
		<meta name="twitter:site" content="@ncs_pl" />
		<meta name="twitter:title" content="{{ .Title }}" />
		<meta name="twitter:description" content="{{ .Description }}" />
		<meta name="twitter:image" content="{{ .Site }}/assets/crocodile.jpg" />
		<meta property="og:title" content="{{ .Title }}" />
		<meta property="og:type" content="website" />
		<meta property="og:url" content="{{ .Site }}" />
		<meta property="og:description" content="{{ .Description }}" />
		<meta property="og:determiner" content="" />
		<meta property="og:locale" content="en" />
		<meta property="og:site_name" content="Crocc" />
		<meta property="og:image" content="{{ .Site }}/assets/crocodile.jpg" />
		<link rel="stylesheet" href="/style.css">
		<link rel="icon" href="assets/emoji_crocodile.ico" sizes="any">
		<link rel="mask-icon" href="assets/emoji_crocodile.svg" color="155eca">
	</head>
	<body>
		<header>
			<nav>
				<ul>
					<li><strong>Crocc</strong></li>
					<li><a href="/">Home</a></li>
					<li>
						<a href="https://github.com/n1c00o/crocc"
							target="_blank">Source</a>
					</li>
					<li>
						<a href="https://nc0.fr" target="_blank">Nicolas
							Paul</a>
					</li>
				</ul>
			</nav>
		</header>
		<main>
			{{ .Content }}
		</main>
		<footer>
			<hr />

			<p>Copyright (c) 2023, Nicolas Paul</p>
		</footer>
	</body>
</html>
```