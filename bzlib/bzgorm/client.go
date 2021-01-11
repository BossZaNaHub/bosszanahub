package bzgorm

import "gorm.io/gorm"

type Client interface {
	InitMysql() (*gorm.DB, error)
	InitPsql() (*gorm.DB, error)
}
