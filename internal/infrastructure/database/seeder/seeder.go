package seeder

import (
	"log"

	"loan-service/internal/domain/user"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	borrowers := []user.Borrower{
		{ID: 1, Name: "John Doe", Identification: "ID001"},
		{ID: 2, Name: "Jane Smith", Identification: "ID002"},
		{ID: 3, Name: "Alice Johnson", Identification: "ID003"},
	}

	// Sample Investors
	investors := []user.Investor{
		{ID: 1, Name: "Investor A", Email: "investora@example.com"},
		{ID: 2, Name: "Investor B", Email: "investorb@example.com"},
		{ID: 3, Name: "Investor C", Email: "investorc@example.com"},
	}

	// Insert borrowers
	for _, borrower := range borrowers {
		if err := db.Create(&borrower).Error; err != nil {
			log.Printf("could not seed borrower: %v", err)
		}
	}

	// Insert investors
	for _, investor := range investors {
		if err := db.Create(&investor).Error; err != nil {
			log.Printf("could not seed investor: %v", err)
		}
	}

	log.Println("Sample loans seeded successfully.")
}
