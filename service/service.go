package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ttvs-blockchain/firebaseupdate/constants"
	"github.com/ttvs-blockchain/firebaseupdate/models"
	"github.com/ttvs-blockchain/firebaseupdate/storage"
)

type SimpleSYNCservice struct {
	ctx                context.Context
	localCertDAO       *storage.LocalCertDAO
	globalChainInfoDAO *storage.GlobalChainInfoDAO
	fireStoreDAO       *storage.FireStoreDAO
}

func NewSimpleSYNCservice(ctx context.Context,
	localCertDAo *storage.LocalCertDAO,
	globalChainInfoDAO *storage.GlobalChainInfoDAO,
	fireStoreDAO *storage.FireStoreDAO) *SimpleSYNCservice {
	return &SimpleSYNCservice{
		ctx:                ctx,
		localCertDAO:       localCertDAo,
		globalChainInfoDAO: globalChainInfoDAO,
		fireStoreDAO:       fireStoreDAO,
	}
}

func (s *SimpleSYNCservice) Sync() error {
	type splitGlobalInfo struct {
		merkleTreeRoot       string
		globalChainTxHash    string
		globalChainBlockNum  int
		globalChainTimestamp string
	}
	localCertList, err := s.localCertDAO.GetAllLocalCertificates(s.ctx)
	if err != nil {
		return err
	}
	globalChainInfoList, err := s.globalChainInfoDAO.GetAllGlobalChainInfo(s.ctx)
	if err != nil {
		return err
	}
	globalInfoMap := map[string]splitGlobalInfo{}
	for _, globalInfo := range globalChainInfoList {
		splitCertIDList := strings.Split(globalInfo.CertIDList, ",")
		for i := 0; i < len(splitCertIDList); i++ {
			splitCertIDList[i] = strings.Trim(splitCertIDList[i], "[]\",")
		}
		for _, certID := range splitCertIDList {
			globalInfoMap[certID] = splitGlobalInfo{
				merkleTreeRoot:       globalInfo.MerkleTreeRoot,
				globalChainTxHash:    globalInfo.GlobalChainTxHash,
				globalChainBlockNum:  globalInfo.GlobalChainBlockNum,
				globalChainTimestamp: globalInfo.GlobalChainTimestamp,
			}
		}
	}
	var firebaseCertList []*models.FirebaseCertificate
	for _, localCert := range localCertList {
		localChainTime, err := parseChainTimestamp(localCert.LocalChainTimestamp)
		if err != nil {
			return err
		}
		localCert.LocalChainTimestamp = localChainTime
		localCert.IssueTime = parseIssueTime(localCert.IssueTime)
		fmt.Printf("Local Certificate: %v \n", localCert)
		if globalInfo, ok := globalInfoMap[localCert.CertID]; ok {
			globalChainTimestamp, err := parseChainTimestamp(globalInfo.globalChainTimestamp)
			if err != nil {
				return err
			}
			firebaseCertList = append(firebaseCertList,
				models.NewFirebaseCertificate(
					*localCert,
					globalInfo.globalChainTxHash,
					globalChainTimestamp,
					globalInfo.globalChainBlockNum,
					true,
				),
			)
		} else {
			firebaseCertList = append(firebaseCertList, models.NewFirebaseCertificate(
				*localCert,
				"",
				"",
				0,
				false,
			))
		}
	}
	s.fireStoreDAO.DeleteAllCertificates(s.ctx)
	for _, cert := range firebaseCertList {
		fmt.Printf("Uploading Certificate %v \n", cert)
		if err := s.fireStoreDAO.AddNewCertificate(s.ctx, cert); err != nil {
			return err
		}
	}
	return nil
}

func (s *SimpleSYNCservice) SyncWithTime(numMinutes int) {
	ticker := time.NewTicker(time.Minute * time.Duration(numMinutes))
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				s.Sync()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func parseChainTimestamp(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	timestampInt64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", fmt.Errorf("time parsing error: %v", err)
	}
	fmt.Printf("Timestamp: %d(%s)\n", timestampInt64, s)
	unixTime := time.Unix(timestampInt64, 0)
	formattedTime := unixTime.Format(constants.DateTimeFormat)
	fmt.Printf("Formatted time: %s\n", formattedTime)
	return formattedTime, nil
}

func parseIssueTime(s string) string {
	replaceT := strings.Replace(s, "T", " ", -1)
	replaceZ := strings.Replace(replaceT, "Z", "", -1)
	return replaceZ
}
