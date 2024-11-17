package application

import "loan-service/internal/domain/loan"

type LoanUseCase struct {
	loanService loan.ILoanService
}

func NewLoanUseCase(loanService loan.ILoanService) *LoanUseCase {
	return &LoanUseCase{loanService: loanService}
}

func (u *LoanUseCase) CreateLoan(dataLoan *loan.Loan) (*loan.Loan, error) {
	return u.loanService.CreateLoan(dataLoan)
}

func (u *LoanUseCase) GetLoanList(state string) ([]*loan.Loan, error) {
	return u.loanService.GetLoansByState(state)
}

func (u *LoanUseCase) ApproveLoan(loanID int, details loan.ApprovalDetails) (*loan.Loan, error) {
	return u.loanService.ApproveLoan(loanID, details)
}

func (u *LoanUseCase) InvestLoan(loanID int, investorID int, amount float64) (*loan.Loan, error) {
	return u.loanService.InvestInLoan(loanID, investorID, amount)
}

func (u *LoanUseCase) DisburseLoan(loanID int, details loan.DisbursementDetails) (*loan.Loan, error) {
	return u.loanService.DisburseLoan(loanID, details)
}
