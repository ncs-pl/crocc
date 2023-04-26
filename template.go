// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
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
}

// GenerateHTML generates the HTML file from the Markdown document.
func GenerateHTML(file io.Writer, fm FrontMatter, content string) error {
	return htmlTemplate.Execute(file, TemplateData{
		Title:           fm.Title,
		Description:     fm.Description,
		PublicationTime: fm.PublicationTime.Format(time.RFC3339),
		LastUpdateTime:  fm.LastUpdateTime.Format(time.RFC3339),
		Keywords:        fmt.Sprintf("%s", fm.Keywords),
		Author:          fm.Author,
		Content:         content,
		Site:            *url,
		Generator:       fmt.Sprintf("crocc %s (https://crocc.nc0.fr)", version),
	})
}
