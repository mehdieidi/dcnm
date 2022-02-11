package transact

import (
	"fmt"

	"github.com/MehdiEidi/dcnm/core"
)

func NewTransactionLogger(loggerType string) (core.TransactionLogger, error) {
	switch loggerType {
	case "file":
		return NewFileTransactionLogger("./transactions.txt")

	case "postgres":
		params := PostgresDbParams{
			host: "localhost", dbName: "kvs",
			user: "test", password: "hunter2",
		}
		return NewPostgresTransactionLogger(params)

	default:
		return nil, fmt.Errorf("no such transaction logger %s", loggerType)
	}
}
