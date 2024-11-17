package database

import (
	"loan-service/internal/domain/loan"
	"loan-service/internal/domain/user"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Running database migrations...")
	err = db.AutoMigrate(
		&loan.Loan{},
		&loan.ApprovalDetails{},
		&loan.Investment{},
		&loan.DisbursementDetails{},
		&user.Borrower{},
		&user.Investor{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
