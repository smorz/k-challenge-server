// Please set the database constants in this file before running.

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

	route = "/last-trade"
	port  = 8000
)

func main() {

	// setup log
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// setup db
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	// create a handler intrface
	ca, err := challenge.NewLastTradeServer(db)
	if err != nil {
		log.Fatal(err)
	}

	// create a router instace
	router := mux.NewRouter()

	//setup the router
	router.Handle(route, ca)

	//serve
	fmt.Printf("Serve at port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))

}
