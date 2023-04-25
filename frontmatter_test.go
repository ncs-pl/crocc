// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestParseFrontMatter(t *testing.T) {
	tests := []struct {
		input []byte
		want  FrontMatter
		rest  string
	}{
		{
			input: []byte(`---
title: "Test"
description: "Test"
publication_time: 2020-01-01T00:00:00Z
last_update_time: 2020-01-01T00:00:00Z
keywords: ["test"]
author: "Test"
hide: false
---
# Test`),
			want: FrontMatter{
				Title:           "Test",
				Description:     "Test",
				PublicationTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				LastUpdateTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Keywords:        []string{"test"},
				Author:          "Test",
				Hide:            false,
			},
			rest: "# Test",
		},
		{
			input: []byte(`---
title: "Test"
description: "Test"
publication_time: 2020-01-01T00:00:00Z
keywords: ["test", "test2", "test3"]
---
# Test`),
			want: FrontMatter{
				Title:           "Test",
				Description:     "Test",
				PublicationTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Keywords:        []string{"test", "test2", "test3"},
			},
			rest: "# Test",
		},
	}

	for _, tt := range tests {
		fm, rest, err := ParseFrontMatter(tt.input)

		if err != nil {
			t.Fatalf("ParseFrontMatter(%v) got error %v", tt.input, err)
		}

		if !cmp.Equal(fm, tt.want) {
			t.Fatalf("ParseFrontMatter(%v) got %v want %v", tt.input, fm, tt.want)
		}

		if string(rest) != tt.rest {
			t.Fatalf("ParseFrontMatter(%v) got %v want %v", tt.input, rest, tt.rest)
		}
	}
}

func FuzzParseFrontMatter(f *testing.F) {
	// Do not fuzz *_time fields as time.Time is not supported by f.Fuzz().
	f.Fuzz(func(t *testing.T, title string, description string, keyword1 string, keyword2 string, keyword3 string, author string, hide bool, text string) {
		input := []byte(`---
title: "` + title + `"
description: "` + description + `"
publication_time: 2020-01-01T00:00:00Z
last_update_time: 2020-01-01T00:00:00Z
keywords: ["` + keyword1 + `", "` + keyword2 + `", "` + keyword2 + `"]
author: "` + author + `"
hide: ` + strconv.FormatBool(hide) + `
---
` + text)

		fm, rest, err := ParseFrontMatter(input)

		if err != nil {
			t.Fatalf("ParseFrontMatter(%v) got error %v", input, err)
		}

		if string(rest) != text {
			t.Fatalf("ParseFrontMatter(%v) got %v want %v", input, rest, text)
		}

		result := FrontMatter{
			Title:           title,
			Description:     description,
			Keywords:        []string{keyword1, keyword2, keyword3},
			PublicationTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			LastUpdateTime:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Author:          author,
			Hide:            hide,
		}

		if !cmp.Equal(fm, result) {
			t.Fatalf("ParseFrontMatter(%v) got %v want %v", input, fm, result)
		}

	})
}
