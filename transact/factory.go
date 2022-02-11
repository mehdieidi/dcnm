package transact

import (
	"fmt"

	"github.com/MehdiEidi/dcnm/core"
)

func NewTransactionLogger(s string) (core.TransactionLogger, error) {
	switch s {
	case "file":
		return NewFileTransactionLogger("./transactions.txt")

	case "postgres":
		params := PostgresDbParams{
			host: "localhost", dbName: "kvs",
			user: "test", password: "hunter2",
		}
		return NewPostgresTransactionLogger(params)

	default:
		return nil, fmt.Errorf("no such transaction logger %s", s)
	}
}
