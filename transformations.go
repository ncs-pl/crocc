// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// TransformMarkdownFile simply copy a non-markdown file to the output
// directory.
func TransformNonMarkdownFile(i, o string) error {
	input, err := os.ReadFile(i)
	if err != nil {
		return err
	}

	if err := os.WriteFile(o, input, 0666); err != nil {
		return err
	}

	log.Printf("copied file %q to %q", i, o)

	return nil
}

// TransformDirectory creates a directory in the output directory.
func TransformDirectory(o string) error {
	if err := os.MkdirAll(o, 0777); err != nil {
		return err
	}

	log.Printf("created directory %q", o)
	return nil
}

// TransformMarkdownFile generates the corresponding HTML document from a
// Markdown file.
func TransformMarkdownFile(i, o string) error {
	raw, err := os.ReadFile(i)
	if err != nil {
		return err
	}

	// Parse front matter
	fm, md, err := ParseFrontMatter(raw)
	if err != nil {
		return err
	}

	// Skip hidden files unless -hidden is specified
	if fm.Hide && !*generateHidden {
		log.Printf("skipped hidden file %q", i)
		return nil
	}

	// Render Markdown to HTML
	pExtensions := parser.Tables | parser.FencedCode |
		parser.Autolink | parser.Strikethrough | parser.SpaceHeadings |
		parser.HeadingIDs | parser.BackslashLineBreak |
		parser.AutoHeadingIDs | parser.Footnotes | parser.SuperSubscript |
		parser.NoIntraEmphasis
	p := parser.NewWithExtensions(pExtensions)
	ast := p.Parse(md)

	htmlFlags := html.Smartypants | html.SmartypantsFractions |
		html.SmartypantsDashes | html.SmartypantsLatexDashes |
		html.HrefTargetBlank | html.LazyLoadImages
	renderer := html.NewRenderer(html.RendererOptions{Flags: htmlFlags})
	html := markdown.Render(ast, renderer)

	c, err := GenerateHTML(fm, string(html))
	if err != nil {
		return err
	}

	if err := os.WriteFile(o, c, 0666); err != nil {
		return err
	}

	log.Printf("generated file %q", o)
	return nil
}
