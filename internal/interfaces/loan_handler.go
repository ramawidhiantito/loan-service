package interfaces

import (
	"encoding/json"
	"net/http"

	"loan-service/internal/application"
	"loan-service/internal/domain/loan"
)

type LoanHandler struct {
	usecase *application.LoanUseCase
}

func NewLoanHandler(usecase *application.LoanUseCase) *LoanHandler {
	return &LoanHandler{usecase: usecase}
}

func (h *LoanHandler) CreateLoan(w http.ResponseWriter, r *http.Request) {
	var loan loan.Loan
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	data, err := h.usecase.CreateLoan(&loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *LoanHandler) ApproveLoan(w http.ResponseWriter, r *http.Request) {
	var details loan.ApprovalDetails
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	data, err := h.usecase.ApproveLoan(details.LoanID, details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *LoanHandler) GetListLoan(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state == "" {
		http.Error(w, "state is empty", http.StatusBadRequest)
		return
	}
	data, err := h.usecase.GetLoanList(loan.Approved)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *LoanHandler) InvestLoan(w http.ResponseWriter, r *http.Request) {
	var details loan.Investment
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	data, err := h.usecase.InvestLoan(details.LoanID, details.InvestorID, details.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *LoanHandler) DisburseLoan(w http.ResponseWriter, r *http.Request) {
	var details loan.DisbursementDetails
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	data, err := h.usecase.DisburseLoan(details.LoanID, details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
