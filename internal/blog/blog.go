package blog

import (
	"os"
	"path/filepath"

	"github.com/cryring/blog_backend/internal/config"
	"github.com/cryring/blog_backend/internal/log"
	"github.com/cryring/blog_backend/internal/utils"
)

type Blog struct {
	category Category
	filename string
	content  string
}

func New(category Category, filename string) *Blog {
	return &Blog{
		category: category,
		filename: filename,
	}
}

func (b *Blog) Load() error {
	path := b.buildFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		log.Errorf("read blog file[%s] failed: %v", path, err)
		return err
	}
	b.content = utils.Bytes2String(data)
	return nil
}

func (b *Blog) Content() string {
	return b.content
}

func (b *Blog) buildFilePath() string {
	return filepath.Join(config.GetConfig().RootDir, b.category.String(), b.filename)
}
