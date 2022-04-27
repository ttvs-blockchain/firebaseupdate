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
		MerkleTreePath:      "",
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

func MakeStageTwoDummyData(dao *FirebaseCertificateDAO) error {
	ctx := context.Background()
	dao.DeleteAllCertificates(ctx)
	cert1 := models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "001",
		PersonID:            "001",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           1,
		IssueTime:           "2022-01-22 00:00:00",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "7886781e6ab3770f1abb7859b6e0078d0c77cbfd16df2ce6f50649d1ef560203",
		LocalChainBlockNum:  1,
		LocalChainTimestamp: "2022-01-22 00:00:00",
	},
		"9931d2f5b3b76b97719461e01500758361f2ae93ff46f58e2ce378bbb4cc7bda",
		"2022-01-22 00:00:00",
		1,
		true)
	cert2 := models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "002",
		PersonID:            "001",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           2,
		IssueTime:           "2022-02-22 00:00:00",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "5d711b89c2b238adcbaec0564d81855ba492badee848b70b3203ab8ffe43a270",
		LocalChainBlockNum:  145,
		LocalChainTimestamp: "2022-02-22 00:00:00",
	},
		"3d0e655bb36bd111e427bac5485d4f7c882074541ca54fb026f7424d46b8d10f",
		"2022-02-22 00:00:00",
		96,
		false)
	if err := dao.AddNewCertificate(ctx, cert1); err != nil {
		return err
	}
	if err := dao.AddNewCertificate(ctx, cert2); err != nil {
		return err
	}
	return nil
}

func MakeStageThreeDummyData(dao *FirebaseCertificateDAO) error {
	ctx := context.Background()
	dao.DeleteAllCertificates(ctx)
	cert1 := models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "001",
		PersonID:            "001",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           1,
		IssueTime:           "2022-01-22 00:00:00",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "7886781e6ab3770f1abb7859b6e0078d0c77cbfd16df2ce6f50649d1ef560203",
		LocalChainBlockNum:  1,
		LocalChainTimestamp: "2022-01-22 00:00:00",
	},
		"9931d2f5b3b76b97719461e01500758361f2ae93ff46f58e2ce378bbb4cc7bda",
		"2022-01-22 00:00:00",
		55,
		true)
	cert2 := models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "002",
		PersonID:            "001",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           2,
		IssueTime:           "2022-02-22 00:00:00",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "5d711b89c2b238adcbaec0564d81855ba492badee848b70b3203ab8ffe43a270",
		LocalChainBlockNum:  145,
		LocalChainTimestamp: "2022-02-22 00:00:00",
	},
		"3d0e655bb36bd111e427bac5485d4f7c882074541ca54fb026f7424d46b8d10f",
		"2022-02-22 00:00:00",
		96,
		true)
	cert3 := models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "003",
		PersonID:            "001",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           3,
		IssueTime:           "2022-03-22 00:00:00",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "c380779f6175766fdbe90940851fff3995d343c63bbb82f816843c1d5100865e",
		LocalChainBlockNum:  300,
		LocalChainTimestamp: "2022-03-22 00:00:00",
	},
		"c61b60be803805105bacb648b03a8f5aee7ccc34fd1eb280dae3db0d4156a753",
		"2022-03-22 00:00:00",
		200,
		false)
	if err := dao.AddNewCertificate(ctx, cert1); err != nil {
		return err
	}
	if err := dao.AddNewCertificate(ctx, cert2); err != nil {
		return err
	}
	if err := dao.AddNewCertificate(ctx, cert3); err != nil {
		return err
	}
	return nil
}