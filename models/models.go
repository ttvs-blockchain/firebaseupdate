package models

const (
	TableNameLocalCertificate = "localCertificate"
	TableNameGlobalChainInfo  = "globalChainInfo"
)

// LocalCertificate is the database model for the table localCertificate
type LocalCertificate struct {
	ID                  int    `gorm:"column:id"`
	CertID              string `gorm:"column:certID"`
	PersonID            string `gorm:"column:personID"`
	Name                string `gorm:"column:name"`
	Brand               string `gorm:"column:brand"`
	NumOfDose           int    `gorm:"column:numOfDose"`
	IssueTime           string `gorm:"column:issueTime"`
	Issuer              string `gorm:"column:issuer"`
	Remark              string `gorm:"column:remark"`
	PersonHash          string `gorm:"column:personHash"`
	MerkleTreePath      string `gorm:"column:merkleTreePath"`
	MerkleTreeIndexes   string `gorm:"column:merkleTreeIndexes"`
	GlobalRootID        string `gorm:"column:globalRootID"`
	LocalChainID        string `gorm:"column:localChainID"`
	LocalChainTxHash    string `gorm:"column:localChainTxHash"`
	LocalChainBlockNum  int    `gorm:"column:localChainBlockNum"`
	LocalChainTimestamp string `gorm:"column:localChainTimeStamp"`
}

func (*LocalCertificate) TableName() string {
	return TableNameLocalCertificate
}

// GlobalChainInfo is the database model for the table globalChainInfo
type GlobalChainInfo struct {
	ID                   int
	CertIDList           string `gorm:"column:certIDList"`
	MerkleTreeRoot       string `gorm:"column:merkleTreeRoot"`
	GlobalChainTxHash    string `gorm:"column:globalChainTxHash"`
	GlobalChainBlockNum  int    `gorm:"column:globalChainBlockNum"`
	GlobalChainTimestamp string `gorm:"column:globalChainTimeStamp"`
}

func (*GlobalChainInfo) TableName() string {
	return TableNameGlobalChainInfo
}

type FirebaseCertificate struct {
	data map[string]interface{}
}

func (f *FirebaseCertificate) Data() map[string]interface{} {
	return f.data
}

func NewFirebaseCertificate(localCert LocalCertificate, globalTxHash, globalTimestamp string,
	gobalBlockNum int, isValidated bool) *FirebaseCertificate {
	return &FirebaseCertificate{data: map[string]interface{}{
		"cert_id":                localCert.CertID,
		"person_id":              localCert.PersonID,
		"name":                   localCert.Name,
		"brand":                  localCert.Brand,
		"num_dose":               localCert.NumOfDose,
		"issue_time":             localCert.IssueTime,
		"issuer":                 localCert.Issuer,
		"remark":                 localCert.Remark,
		"global_chain_tx_hash":   globalTxHash,
		"global_chain_block_num": gobalBlockNum,
		"global_chain_timestamp": globalTimestamp,
		"local_chain_id":         localCert.LocalChainID,
		"local_chain_tx_hash":    localCert.LocalChainTxHash,
		"local_chain_block_num":  localCert.LocalChainBlockNum,
		"local_chain_timestamp":  localCert.LocalChainTimestamp,
		"is_validated":           isValidated,
		"merkle_tree_path":       localCert.MerkleTreePath,
		"merkle_tree_indexes":    localCert.MerkleTreeIndexes,
		"global_root_id":         localCert.GlobalRootID,
	}}
}
