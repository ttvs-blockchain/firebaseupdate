package storage

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/ttvs-blockchain/firebaseupdate/constants"
	"github.com/ttvs-blockchain/firebaseupdate/models"
	"google.golang.org/api/iterator"
)

type FireStoreDAO struct {
	client *firestore.Client
}

func NewFireStoreDAO(client *firestore.Client) *FireStoreDAO {
	return &FireStoreDAO{client: client}
}

func (f *FireStoreDAO) AddNewCertificate(ctx context.Context,
	cert *models.FirebaseCertificate) error {
	_, _, err := f.client.Collection(constants.FirebaseCollectionName).Add(ctx, cert.Data())
	return err
}

func (f *FireStoreDAO) DeleteAllCertificates(ctx context.Context) error {
	for {
		// Get a batch of documents
		ref := f.client.Collection(constants.FirebaseCollectionName)
		iter := ref.Limit(constants.FirestoreDeleteBatchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := f.client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

func (f *FireStoreDAO) Deinit() error {
	err := f.client.Close()
	return err
}
