package mysql

import (
	"fmt"

	"github.com/cryring/blog_backend/internal/blog"
)

func (d *Dao) AutoMigrate() error {
	if err := d.db.Migrator().AutoMigrate(&Tag{}); err != nil && !IsAlreadyExistErr(err) {
		return fmt.Errorf("auto migrate db table[tag] failed: %v", err)
	}
	if err := d.db.Migrator().AutoMigrate(&TagMap{}); err != nil && !IsAlreadyExistErr(err) {
		return fmt.Errorf("auto migrate db table[tagmap] failed: %v", err)
	}

	categories := []blog.Category{blog.Golang, blog.Cpp, blog.Rust}
	for _, cg := range categories {
		table := &Blog{Category: cg}
		err := d.db.Set("gorm:table_options",
			"ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_bin",
		).Table(table.TableName()).AutoMigrate(table)
		if err != nil && !IsAlreadyExistErr(err) {
			return fmt.Errorf("auto migrate db table[%s] failed: %v", table.TableName(), err)
		}
	}
	return nil
}
