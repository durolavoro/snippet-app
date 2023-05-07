package model

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("no matching record found")

type CreateSnippetsPayload struct {
	Snippets []CreateSnippetPayload `json:"snippets"`
}

type CreateSnippetPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Expires int    `json:"expires"`
}

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
