package loan

import "time"

const (
	Proposed  = "proposed"
	Approved  = "approved"
	Invested  = "invested"
	Disbursed = "disbursed"
)

type Loan struct {
	ID              int                  `gorm:"primaryKey" json:"id"`
	BorrowerID      string               `json:"borrower_id"`
	PrincipalAmount float64              `json:"principal_amount"`
	Rate            float64              `json:"rate"`
	ROI             float64              `json:"roi"`
	State           string               `json:"state"`
	ApprovalDetails *ApprovalDetails     `gorm:"foreignKey:LoanID" json:"approval_details,omitempty"`
	Investments     []Investment         `gorm:"foreignKey:LoanID" json:"investments,omitempty"`
	Disbursement    *DisbursementDetails `gorm:"foreignKey:LoanID" json:"disbursement,omitempty"`
	AgreementLetter string               `json:"agreement_letter"`
	TotalInvested   float64              `json:"total_invested"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
}

type ApprovalDetails struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	LoanID       int       `gorm:"index" json:"loan_id"`
	EmployeeID   int       `json:"employee_id"`
	ApprovalDate time.Time `json:"approval_date"`
	PictureURL   string    `json:"picture_url"`
}

type Investment struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	LoanID     int     `gorm:"index" json:"loan_id"`
	InvestorID int     `json:"investor_id"`
	Amount     float64 `json:"amount"`
}

type DisbursementDetails struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	LoanID           int       `gorm:"index" json:"loan_id"`
	EmployeeID       int       `json:"employee_id"`
	DisbursementDate time.Time `json:"disbursement_date"`
	AgreementFileURL string    `json:"agreement_file_url"`
}
