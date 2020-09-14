package gorm_demo

import (
	//"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
添加条件 and
*/
func AddCondition(column string, operator string, value interface{}) func(db *gorm.DB) *gorm.DB {
	if value != nil && value != "" {
		switch operator {
		case "=":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" = ? ", value)
			}
		case "like":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" like ? ", "%"+strings.Replace(value.(string), "_", "\\_", -1)+"%")
			}
		case "in":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" in ( ? ) ", value)
			}
		case "!=":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" != ? ", value)
			}
		case ">":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" > ? ", value)
			}
		case "<":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" < ? ", value)
			}
		case ">=":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" >= ? ", value)
			}
		case "<=":
			return func(db *gorm.DB) *gorm.DB {
				return db.Where(column+" <= ?", value)
			}
		}

	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
