package models

import (
	"errors"
	"time"
)

// ErrNoRecord CURD error string
var ErrNoRecord = errors.New("models: no matching record found ss")

// Snippet model
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
