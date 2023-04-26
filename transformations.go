// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// TransformMarkdownFile simply copy a non-markdown file to the output
// directory.
func TransformNonMarkdownFile(inputDir, inputFile, outputDir string) error {
	inputPath := filepath.Join(inputDir, inputFile)

	input, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	outputPath := filepath.Join(outputDir, inputFile)
	if err := os.WriteFile(outputPath, input, 0666); err != nil {
		return err
	}

	log.Printf("copied file %q to %q", inputPath, outputPath)

	return nil
}

// TransformDirectory creates a directory in the output directory.
func TransformDirectory(inputDir, inputFile, outputDir string) error {
	outputPath := filepath.Join(outputDir, inputFile)

	if err := os.MkdirAll(outputPath, 0777); err != nil {
		return err
	}

	log.Printf("created directory %q", outputPath)

	return nil
}

// TransformMarkdownFile generates the corresponding HTML document from a
// Markdown file.
func TransformMarkdownFile(inputDir, inputFile, outputDir string) error {
	inputPath := filepath.Join(inputDir, inputFile)

	// The output file is the same as the input file, but with a different
	// extension.
	fn := strings.TrimSuffix(inputFile, filepath.Ext(inputFile)) + ".html"
	outputPath := filepath.Join(outputDir, fn)

	contentRaw, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Parse front matter
	fm, contentMD, err := ParseFrontMatter(contentRaw)
	if err != nil {
		return err
	}

	// Skip hidden files unless -hidden is specified
	if fm.Hide && !*generateHidden {
		log.Printf("skipped hidden file %q", inputPath)
		return nil
	}

	// Render Markdown to HTML
	pExtensions := parser.Tables | parser.FencedCode |
		parser.Autolink | parser.Strikethrough | parser.SpaceHeadings |
		parser.HeadingIDs | parser.BackslashLineBreak |
		parser.AutoHeadingIDs | parser.Footnotes | parser.SuperSubscript |
		parser.NoIntraEmphasis
	p := parser.NewWithExtensions(pExtensions)
	doc := p.Parse(contentMD)

	htmlFlags := html.Smartypants | html.SmartypantsFractions |
		html.SmartypantsDashes | html.SmartypantsLatexDashes |
		html.HrefTargetBlank | html.LazyLoadImages
	renderer := html.NewRenderer(html.RendererOptions{Flags: htmlFlags})
	contentHTML := markdown.Render(doc, renderer)

	c, err := GenerateHTML(fm, string(contentHTML))
	if err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, c, 0666); err != nil {
		return err
	}

	log.Printf("generated file %q", outputPath)
	return nil
}
