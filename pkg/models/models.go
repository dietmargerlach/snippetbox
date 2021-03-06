package models

import (
	"errors"
	"time"
)

// ErrNoRecord exported
var ErrNoRecord = errors.New("models: no matching record found")

// Snippet exported
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
