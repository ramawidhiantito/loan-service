package loan

import (
	"fmt"
	"time"
)

type ILoanService interface {
	CreateLoan(dataLoan *Loan) (*Loan, error)
	GetLoansByState(state string) ([]*Loan, error)
	ApproveLoan(loanID int, details ApprovalDetails) (*Loan, error)
	DisburseLoan(loanID int, details DisbursementDetails) (*Loan, error)
	InvestInLoan(loanID int, investorID int, amount float64) (*Loan, error)
}

type LoanService struct {
	repo ILoanRepository
}

func NewLoanService(repository ILoanRepository) ILoanService {
	return &LoanService{repo: repository}
}

func (s *LoanService) CreateLoan(dataLoan *Loan) (*Loan, error) {
	dataLoan.State = Proposed
	err := s.repo.Save(dataLoan)
	if err != nil {
		return nil, err
	}

	return dataLoan, nil
}

func (s *LoanService) GetLoansByState(state string) ([]*Loan, error) {
	switch state {
	case Proposed, Approved, Invested, Disbursed:
		return s.repo.GetAllLoanByState(state)
	default:
		return nil, fmt.Errorf("error: unknown state")
	}
}

func (s *LoanService) ApproveLoan(loanID int, details ApprovalDetails) (*Loan, error) {
	loan, err := s.repo.GetByID(loanID)
	if err != nil {
		return nil, err
	}

	if loan.State != Proposed {
		return nil, fmt.Errorf("loan is already in state: %s", loan.State)
	}

	loan.State = Approved
	details.ApprovalDate = time.Now()
	loan.ApprovalDetails = &details

	err = s.repo.Save(loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *LoanService) DisburseLoan(loanID int, details DisbursementDetails) (*Loan, error) {
	loan, err := s.repo.GetByID(loanID)
	if err != nil {
		return nil, err
	}

	if loan.State != Invested {
		return nil, fmt.Errorf("loan is not invested yet")
	}

	loan.State = Disbursed
	details.DisbursementDate = time.Now()
	loan.AgreementLetter = details.AgreementFileURL
	loan.Disbursement = &details

	err = s.repo.Save(loan)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (s *LoanService) InvestInLoan(loanID int, investorID int, amount float64) (*Loan, error) {
	err := s.repo.InvestInLoan(loanID, investorID, amount)
	if err != nil {
		return nil, fmt.Errorf("could not invest in loan: %v", err)
	}

	data, err := s.repo.GetByID(loanID)
	if err != nil {
		return nil, fmt.Errorf("loan id not found: %v", err)
	}
	return data, nil
}
