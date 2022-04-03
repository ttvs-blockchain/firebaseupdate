package storage

import (
	"context"

	"github.com/ttvs-blockchain/firebaseupdate/models"
)

func MakeStageOneDummyData(dao *FirebaseCertificateDAO) error {
	ctx := context.Background()
	dao.DeleteAllCertificates(ctx)
	cert := models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "001",
		PersonID:            "001",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           1,
		IssueTime:           "2022-01-22 00:00:00",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkelTreePath:      "",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "7886781e6ab3770f1abb7859b6e0078d0c77cbfd16df2ce6f50649d1ef560203",
		LocalChainBlockNum:  1,
		LocalChainTimestamp: "2022-01-22 00:00:00",
	},
		"9931d2f5b3b76b97719461e01500758361f2ae93ff46f58e2ce378bbb4cc7bda",
		"2022-01-22 00:00:00",
		1,
		false)
	if err := dao.AddNewCertificate(ctx, cert); err != nil {
		return err
	}
	return nil
}

func MakeStageTwoDummyData()
