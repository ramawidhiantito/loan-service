package user

import (
	"errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetBorrowerByID(id int) (*Borrower, error)
	GetInvestorByID(id int) (*Investor, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewLoanRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) GetBorrowerByID(id int) (*Borrower, error) {
	var borrower Borrower
	if err := u.DB.Take(&borrower, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &borrower, nil
}

func (u *UserRepository) GetInvestorByID(id int) (*Investor, error) {
	var investor Investor
	if err := u.DB.Take(&investor, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &investor, nil
}
