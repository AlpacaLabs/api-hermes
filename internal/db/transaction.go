package db

import (
	"context"
	"database/sql"

	hermesV1 "github.com/AlpacaLabs/protorepo-hermes-go/alpacalabs/hermes/v1"
	"github.com/golang-sql/sqlexp"
	"github.com/golang/protobuf/proto"
)

type Transaction interface {
	SaveSendEmailRequest(ctx context.Context, in hermesV1.SendEmailRequest) error
	SaveSendSmsRequest(ctx context.Context, in hermesV1.SendSmsRequest) error
}

type txImpl struct {
	tx *sql.Tx
}

func (tx txImpl) SaveSendEmailRequest(ctx context.Context, in hermesV1.SendEmailRequest) error {
	var q sqlexp.Querier
	q = tx.tx

	b, err := proto.Marshal(&in)
	if err != nil {
		return err
	}

	query := `
INSERT INTO send_email_request(event_id, payload) 
 VALUES($1, $2)
`

	// TODO extract event ID from context
	eventID := "eventID"

	_, err = q.ExecContext(ctx, query, eventID, b)

	return err
}

func (tx txImpl) SaveSendSmsRequest(ctx context.Context, in hermesV1.SendSmsRequest) error {
	var q sqlexp.Querier
	q = tx.tx

	b, err := proto.Marshal(&in)
	if err != nil {
		return err
	}

	query := `
INSERT INTO send_email_request(event_id, payload) 
 VALUES($1, $2)
`

	// TODO extract event ID from context
	eventID := "eventID"

	_, err = q.ExecContext(ctx, query, eventID, b)

	return err
}
