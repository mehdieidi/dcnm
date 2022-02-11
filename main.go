package main

import (
	"log"
	"os"

	"github.com/MehdiEidi/dcnm/core"
	"github.com/MehdiEidi/dcnm/frontend"
	"github.com/MehdiEidi/dcnm/transact"
)

func main() {
	// Creating TransactionLogger. An adapter that will plug into the core application's TransactionLogger plug.
	tl, err := transact.NewTransactionLogger(os.Getenv("TLOG_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	// Creating Core and telling it which TransactionLogger to use. This is a "driven agent".
	store := core.NewKeyValueStore().WithTransactionLogger(tl)
	store.Restore()

	// Creating the frontend. This is a "driving agent".
	fe, err := frontend.NewFrontEnd(os.Getenv("FRONTEND_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(fe.Start(store))
}
