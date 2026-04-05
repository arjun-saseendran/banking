package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arjun-saseendran/banking/domain"
	"github.com/arjun-saseendran/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Envoriment variable not defined!")
	}
}

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[1-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[1-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[1-9]+}/account/{account_id:[1-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
