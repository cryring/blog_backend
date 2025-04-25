package mysql

import (
	"fmt"
	"time"

	"github.com/cryring/blog_backend/internal/blog"
)

type Blog struct {
	ID         int           `gorm:"COLUMN:id;type:int;primaryKey;not null;"      json:"id"`
	Category   blog.Category `gorm:"-"                                            json:"type"`
	Title      string        `gorm:"COLUMN:title;type:varchar(256);not null;"     json:"title"`
	FileName   string        `gorm:"COLUMN:filename;type:varchar(512);not null;"  json:"filename"`
	Preview    string        `gorm:"COLUMN:preview;type:varchar(512);not null;"   json:"preview"`
	Tags       string        `gorm:"COLUMN:tags;type:varchar(64);"                json:"tags"`
	CreateTime time.Time     `gorm:"COLUMN:create_time;not null;"                 json:"create_time"`
	UpdateTime time.Time     `gorm:"COLUMN:update_time;not null;"                 json:"update_time"`
}

func (b *Blog) TableName() string {
	return fmt.Sprintf("%v_blogs", b.Category)
}

type Tag struct {
	ID    int    `gorm:"COLUMN:id;type:int;primaryKey;not null;" json:"id"`
	Name  string `gorm:"COLUMN:name;type:varchar(64);not null;"  json:"name"`
	Count int    `gorm:"COLUMN:id;type:int;not null;"            json:"count"`
}

func (Tag) TableName() string {
	return "tags"
}

type TagMap struct {
	TagID  int `gorm:"COLUMN:tag_id;type:int;primaryKey;not null;"    json:"tag_id"`
	BlogID int `gorm:"COLUMN:blog_id;type:int;primaryKey; not null;"  json:"blog_id"`
}

func (TagMap) TableName() string {
	return "tagmap"
}
