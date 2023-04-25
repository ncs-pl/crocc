// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"time"

	"github.com/adrg/frontmatter"
)

// FrontMatter is the front matter of a Markdown document.
type FrontMatter struct {
	Title           string    `yaml:"title"`
	Description     string    `yaml:"description"`
	PublicationTime time.Time `yaml:"publication_time"`
	LastUpdateTime  time.Time `yaml:"last_update_time"`
	Keywords        []string  `yaml:"keywords"`
	Author          string    `yaml:"author"`
	Hide            bool      `yaml:"hide"`
}

// ParseFrontMatter parses the front matter of a Markdown document.
// It returns a FrontMatter struct and the Markdown document without the
// front matter.
func ParseFrontMatter(r []byte) (FrontMatter, []byte, error) {
	var fm FrontMatter

	reader := bytes.NewReader(r)

	rest, err := frontmatter.MustParse(reader, &fm)
	if err != nil {
		return FrontMatter{}, nil, err
	}

	return fm, rest, nil
}
