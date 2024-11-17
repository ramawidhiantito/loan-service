package main

import (
	"log"
	"net/http"

	"loan-service/internal/application"
	"loan-service/internal/domain/loan"
	"loan-service/internal/infrastructure/database"
	"loan-service/internal/infrastructure/database/seeder"
	"loan-service/internal/infrastructure/logger"
	"loan-service/internal/interfaces"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	logger.Init()

	// Load configuration using Viper
	if err := loadConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := viper.GetString("database.dsn")
	db, err := database.NewDB(dsn)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	seeder.Seeder(db)

	// Initialize the repository and loan service
	loanRepository := loan.NewLoanRepository(db)
	loanService := loan.NewLoanService(loanRepository)

	// Initialize the use case, injecting the service and Kafka producer
	loanUseCase := application.NewLoanUseCase(loanService)

	// Initialize HTTP router and handlers
	router := mux.NewRouter()
	loanHandler := interfaces.NewLoanHandler(loanUseCase)

	// Set up routes
	router.HandleFunc("/loan/create", loanHandler.CreateLoan).Methods("POST")
	router.HandleFunc("/loan/approve", loanHandler.ApproveLoan).Methods("POST")
	router.HandleFunc("/loan/list", loanHandler.GetListLoan).Methods("GET")
	router.HandleFunc("/loan/invest", loanHandler.InvestLoan).Methods("POST")
	router.HandleFunc("/loan/disburse", loanHandler.DisburseLoan).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func loadConfig() error {
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Optional: Set defaults if not provided in config
	viper.SetDefault("database.dsn", "host=localhost user=postgres password=postgres123 dbname=loan_service port=5432 sslmode=disable")

	return nil
}
