package user

type Borrower struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"column:name"`
	Identification string `gorm:"column:identification"`
}
