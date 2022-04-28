package storage

import (
	"context"

	"github.com/ttvs-blockchain/firebaseupdate/models"
	"gorm.io/gorm"
)

type GlobalChainInfoDAO struct {
	db *gorm.DB
}

func NewGlobalChainInfoDAO(db *gorm.DB) *GlobalChainInfoDAO {
	return &GlobalChainInfoDAO{
		db: db,
	}
}

func (g *GlobalChainInfoDAO) GetAllGlobalChainInfo(ctx context.Context) ([]*models.GlobalChainInfo, error) {
	var globalChainInfoList []*models.GlobalChainInfo
	result := g.db.WithContext(ctx).Find(&globalChainInfoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return globalChainInfoList, nil
}
