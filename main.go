package main

import (
	"log"
	"net/http"

	"loan-service/internal/application"
	"loan-service/internal/domain/loan"
	"loan-service/internal/infrastructure/database"
	"loan-service/internal/infrastructure/database/seeder"
	"loan-service/internal/infrastructure/kafka"
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

	// Kafka brokers and topic
	brokers := []string{"localhost:9092"}
	topic := "loan-invested"

	if err := kafka.CreateTopic(brokers, topic, 1, 1); err != nil {
		log.Printf("Error creating topic '%s': %v", topic, err)
	}

	kafkaProducer := kafka.NewKafkaProducer(brokers, topic)
	defer kafkaProducer.Close()

	consumer := kafka.NewKafkaConsumer(brokers, topic, "")

	go consumer.ConsumeMessages()

	//init repo,service and usecase
	loanRepository := loan.NewLoanRepository(db, kafkaProducer)
	loanService := loan.NewLoanService(loanRepository)
	loanUseCase := application.NewLoanUseCase(loanService)

	// Router
	router := mux.NewRouter()
	loanHandler := interfaces.NewLoanHandler(loanUseCase)

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
