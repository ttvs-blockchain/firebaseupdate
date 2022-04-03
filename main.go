package main

// import (
// 	"fmt"

// 	"github.com/ttvs-blockchain/firebaseupdate/models"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// func main() {
// 	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect to the database" + err.Error())
// 	}
// 	var globalChainInfoList []*models.GlobalChainInfo
// 	db.Find(&globalChainInfoList)
// 	for _, record := range globalChainInfoList {
// 		fmt.Printf("%+v\n", record)
// 	}
// }

import (
	"context"
	"flag"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/ttvs-blockchain/firebaseupdate/constants"
	"github.com/ttvs-blockchain/firebaseupdate/storage"

	"google.golang.org/api/option"
)

func main() {
	num := flag.Int("n", 5, "# of iterations")
	flag.Parse()
	opt := option.WithCredentialsFile(constants.ServiceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("error initializing app: " + err.Error())
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	dao := storage.NewFirebaseCertificateDAO(client)
	dao.DeleteAllCertificates(context.Background())
	fmt.Printf("Making stage %d dummy data\n", *num)
	if *num == 1 {
		storage.MakeStageOneDummyData(dao)
	} else if *num == 2 {
		storage.MakeStageTwoDummyData(dao)
	} else if *num == 3 {
		storage.MakeStageThreeDummyData(dao)
	}
}
