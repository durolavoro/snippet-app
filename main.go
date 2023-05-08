package main

import (
	"database/sql"
	"github.com/durolavoro/snippet-app/dal"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func openDB(creds string) (*sql.DB, error) {
	db, err := sql.Open("mysql", creds)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	infoLogger := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	// load creds from config
	db, err := openDB("root@/snippetsdb?parseTime=true")
	if err != nil {
		errorLogger.Fatal(err.Error())
	}
	defer db.Close()
	dal := dal.CreateDAL(db)
	// MB dependency injection pattern reference
	a := &App{
		InfoLogger:  infoLogger,
		ErrorLogger: errorLogger,
		DAL:         dal,
	}

	mux := a.routes()

	a.InfoLogger.Println("Server initialised on :4000")
	err = http.ListenAndServe(":4000", mux)
	a.ErrorLogger.Fatalf("Server init failed with err %s", err.Error())
}
