package loan

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ILoanRepository interface {
	GetByID(loanID int) (*Loan, error)
	GetAllLoanByState(state string) ([]*Loan, error)
	Save(loan *Loan) error
	InvestInLoan(loanID int, investorID int, amount float64) error
}

type LoanRepository struct {
	DB *gorm.DB
}

func NewLoanRepository(db *gorm.DB) ILoanRepository {
	return &LoanRepository{DB: db}
}

func (r *LoanRepository) GetByID(loanID int) (*Loan, error) {
	var loan Loan
	if err := r.DB.Preload("Investments").
		Preload("ApprovalDetails").
		Preload("Disbursement").
		First(&loan, "id = ?", loanID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("loan not found")
		}
		return nil, err
	}
	return &loan, nil
}

func (r *LoanRepository) GetAllLoanByState(state string) ([]*Loan, error) {
	var loans []*Loan

	err := r.DB.Model(&Loan{}).Where("state = ?", state).Find(&loans).Error
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (r *LoanRepository) Save(loan *Loan) error {
	return r.DB.Save(loan).Error
}

func (r *LoanRepository) InvestInLoan(loanID int, investorID int, amount float64) error {
	tx := r.DB.Begin()
	defer tx.Rollback()

	// Lock the loan row for the duration of the transaction
	var loan Loan
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&loan, "id = ?", loanID).Error; err != nil {
		return fmt.Errorf("loan not found or could not lock loan: %v", err)
	}

	// Check if the total invested amount exceeds the principal
	if loan.TotalInvested+amount > loan.PrincipalAmount {
		return fmt.Errorf("investment exceeds loan principal amount")
	}

	// Create a new investment record
	investment := Investment{
		LoanID:     loanID,
		InvestorID: investorID,
		Amount:     amount,
	}
	if err := tx.Create(&investment).Error; err != nil {
		return fmt.Errorf("could not add investment: %v", err)
	}

	// Update the loan's total invested amount
	loan.TotalInvested += amount

	if loan.TotalInvested == loan.PrincipalAmount {
		loan.State = Invested
	}
	if err := tx.Save(&loan).Error; err != nil {
		return fmt.Errorf("could not update loan total invested: %v", err)
	}

	tx.Commit()

	return nil
}
