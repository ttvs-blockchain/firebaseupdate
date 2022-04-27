package storage

import "gorm.io/gorm"

type GlobalChainInfoDAO struct {
	db *gorm.DB
}

func NewGlobalChainInfoDAO(db *gorm.DB) *GlobalChainInfoDAO {
	return &GlobalChainInfoDAO{
		db: db,
	}
}
