// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// TemplateData is the data passed to the HTML template.
type TemplateData struct {
	Title           string
	Description     string
	PublicationTime string
	LastUpdateTime  string
	Keywords        string
	Author          string
	Content         string
	Site            string
	Generator       string
	Slug            string
}

// GenerateHTML generates the HTML file from the Markdown document.
func GenerateHTML(fm FrontMatter, slug, content string) ([]byte, error) {
	var buffer bytes.Buffer

	err := htmlTemplate.Execute(&buffer, TemplateData{
		Title:           fm.Title,
		Description:     fm.Description,
		PublicationTime: fm.PublicationTime.Format(time.RFC3339),
		LastUpdateTime:  fm.LastUpdateTime.Format(time.RFC3339),
		Keywords:        strings.Join(fm.Keywords, ","),
		Author:          fm.Author,
		Content:         content,
		Site:            *url,
		Generator:       fmt.Sprintf("crocc %s (https://crocc.nc0.fr)", version),
		Slug:            slug,
	})

	return buffer.Bytes(), err
}
