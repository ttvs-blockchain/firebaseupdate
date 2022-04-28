package storage

import (
	"context"

	"github.com/ttvs-blockchain/firebaseupdate/models"
)

var (
	dummyCert1 *models.FirebaseCertificate = models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "ab60838f-be5b-46ea-8fc5-d5c346a33f9c",
		PersonID:            "ArTEXTPVgqVUU8rtNR1uFTSDwGJ3",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           1,
		IssueTime:           "2022-01-22 13:20:11",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "test",
		MerkleTreeIndexes:   "test",
		GlobalRootID:        "test",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "7886781e6ab3770f1abb7859b6e0078d0c77cbfd16df2ce6f50649d1ef560203",
		LocalChainBlockNum:  1,
		LocalChainTimestamp: "2022-01-22 13:20:11",
	},
		"9931d2f5b3b76b97719461e01500758361f2ae93ff46f58e2ce378bbb4cc7bda",
		"2022-01-23 00:00:00",
		55,
		true)
	dummyCert2 *models.FirebaseCertificate = models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "218e0af8-7b93-4a31-a3f3-7aad053c7b88",
		PersonID:            "ArTEXTPVgqVUU8rtNR1uFTSDwGJ3",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           2,
		IssueTime:           "2022-02-22 08:10:05",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "test",
		MerkleTreeIndexes:   "test",
		GlobalRootID:        "test",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "5d711b89c2b238adcbaec0564d81855ba492badee848b70b3203ab8ffe43a270",
		LocalChainBlockNum:  1,
		LocalChainTimestamp: "2022-02-22 08:10:05",
	},
		"3d0e655bb36bd111e427bac5485d4f7c882074541ca54fb026f7424d46b8d10f",
		"2022-02-23 00:00:00",
		96,
		true)
	dummyCert3 *models.FirebaseCertificate = models.NewFirebaseCertificate(models.LocalCertificate{
		CertID:              "bc51f83b-c4c4-4349-b2f6-93c7faefb4c1",
		PersonID:            "ArTEXTPVgqVUU8rtNR1uFTSDwGJ3",
		Name:                "CoroVac",
		Brand:               "XYZ Co., Ltd.",
		NumOfDose:           3,
		IssueTime:           "2022-04-22 17:32:30",
		Issuer:              "ABC Hospital",
		Remark:              "No remark",
		PersonHash:          "",
		MerkleTreePath:      "test",
		MerkleTreeIndexes:   "test",
		GlobalRootID:        "test",
		LocalChainID:        "local_chain_0",
		LocalChainTxHash:    "c380779f6175766fdbe90940851fff3995d343c63bbb82f816843c1d5100865e",
		LocalChainBlockNum:  1,
		LocalChainTimestamp: "2022-03-22 00:00:00",
	},
		"c61b60be803805105bacb648b03a8f5aee7ccc34fd1eb280dae3db0d4156a753",
		"2022-03-22 00:00:00",
		200,
		false)
)

func MakeStageOneDummyData(dao *FireStoreDAO) error {
	ctx := context.Background()
	dao.DeleteAllCertificates(ctx)

	if err := dao.AddNewCertificate(ctx, dummyCert1); err != nil {
		return err
	}
	return nil
}

func MakeStageTwoDummyData(dao *FireStoreDAO) error {
	ctx := context.Background()
	dao.DeleteAllCertificates(ctx)

	if err := dao.AddNewCertificate(ctx, dummyCert1); err != nil {
		return err
	}
	if err := dao.AddNewCertificate(ctx, dummyCert2); err != nil {
		return err
	}
	return nil
}

func MakeStageThreeDummyData(dao *FireStoreDAO) error {
	ctx := context.Background()
	dao.DeleteAllCertificates(ctx)

	if err := dao.AddNewCertificate(ctx, dummyCert1); err != nil {
		return err
	}
	if err := dao.AddNewCertificate(ctx, dummyCert2); err != nil {
		return err
	}
	if err := dao.AddNewCertificate(ctx, dummyCert3); err != nil {
		return err
	}
	return nil
}
