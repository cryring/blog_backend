package model

import (
	"time"

	"github.com/cryring/blog_backend/internal/blog"
)

type Blog struct {
	ID         int           `json:"id"`
	Category   blog.Category `json:"category"`
	Tags       []string      `json:"tags"`
	Title      string        `json:"title"`
	FileName   string        `json:"filename"`
	Preview    string        `json:"preview"`
	Content    string        `json:"content"`
	CreateTime time.Time     `json:"create_time"`
	UpdateTime time.Time     `json:"update_time"`
}
