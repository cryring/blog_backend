package mysql

import (
	"time"

	"github.com/cryring/blog_backend/internal/blog"
	"github.com/cryring/blog_backend/internal/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dao struct {
	db *gorm.DB
}

func New(dsn string) (*Dao, error) {
	newLogger := logger.New(
		log.NewGormLogger("./log/bridge_sql.log"), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	return &Dao{db: db}, nil
}

func (d *Dao) GetAllBlogs(c blog.Category) ([]*Blog, error) {
	var (
		blogs []*Blog
		m     = &Blog{Category: c}
	)
	db := d.db.Table(m.TableName()).Model(m).Order("id DESC").Find(&blogs)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogs, nil
}

func (d *Dao) GetAllTags() ([]*Tag, error) {
	var tags []*Tag
	db := d.db.Model(&Tag{}).Find(&tags)
	if db.Error != nil {
		return nil, db.Error
	}
	return tags, nil
}

func (d *Dao) GetBlogs(c blog.Category, page, pageSize int) ([]*Blog, error) {
	var (
		blogs []*Blog
		m     = &Blog{Category: c}
	)
	db := d.db.Table(m.TableName()).Model(m).Order("id DESC").Offset(page * pageSize).Limit(pageSize).Find(&blogs)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogs, nil
}

func (d *Dao) GetBlogsByTag(c blog.Category, tag, page, pageSize int) ([]*Blog, error) {
	var ids []int
	db := d.db.Model(&TagMap{}).
		Select("blog_id").
		Where("tag = ?", tag).
		Order("blog_id DESC").
		Offset(page * pageSize).
		Limit(pageSize).
		Find(&ids)
	if db.Error != nil {
		return nil, db.Error
	}

	var blogs []*Blog
	db = d.db.Model(&Blog{Category: c}).Where("id IN ?", ids).Order("id DESC").Find(&blogs)
	if db.Error != nil {
		return nil, db.Error
	}
	return blogs, nil
}
