package storage

import (
	"context"

	"github.com/ttvs-blockchain/firebaseupdate/models"
	"gorm.io/gorm"
)

type LocalCertDAO struct {
	db *gorm.DB
}

func NewLocalCertDAO(db *gorm.DB) *LocalCertDAO {
	return &LocalCertDAO{
		db: db,
	}
}

func (l *LocalCertDAO) GetAllLocalCertificates(ctx context.Context) ([]*models.LocalCertificate, error) {
	var localCertificateList []*models.LocalCertificate
	result := l.db.WithContext(ctx).Find(&localCertificateList)
	if result.Error != nil {
		return nil, result.Error
	}
	return localCertificateList, nil
}
