package transactor

import (
	"context"
	"fmt"

	port "github.com/danisbagus/golang-hexagon-mongo/core/port/transactor"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type (
	Transactor struct {
		client *mongo.Client
	}

	TxKey struct{}
)

func New(client *mongo.Client) port.Transactor {
	return &Transactor{
		client,
	}
}

func (t Transactor) WithinTransaction(tFunc func(ctx context.Context) error) error {
	// define the read and write concerns
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Majority()

	// define transaction options
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	// start session
	session, err := t.client.StartSession()
	if err != nil {
		return fmt.Errorf("%s: %v", "Failed start session mongodb", err)
	}

	// end session at the end
	defer session.EndSession(context.Background())

	// define callback function
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		return nil, tFunc(sessionContext)
	}

	// start transaction
	_, err = session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		return err
	}
	return nil
}
