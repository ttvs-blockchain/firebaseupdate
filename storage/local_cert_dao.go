package storage

import "gorm.io/gorm"

type LocalCertDAO struct {
	db *gorm.DB
}

func NewLocalCertDAO(db *gorm.DB) *LocalCertDAO {
	return &LocalCertDAO{
		db: db,
	}
}
