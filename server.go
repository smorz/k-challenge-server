package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/smorz/k-challenge-server/challenge"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "ktest"

	Route = "/last-trade/"
	Port  = 8000
)

func main() {

	//setup db
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	ca := challenge.NewLastTradeServer(db)
	router.Handle(Route, ca)
	fmt.Printf("Serve at %d port\n", Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(Port), router))

}
