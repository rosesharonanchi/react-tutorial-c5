package piggyservice

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"piggy.com/internal/db/repo"
	"piggy.com/internal/db/sqlc"
	"piggy.com/internal/models"
)

type Service struct {
	repo repo.Repository
}

func NewService(repo repo.Repository) *Service {
	return &Service{repo: repo}
}

// define service methods here

func (s *Service) CreateTransaction(ctx context.Context, userID *int32, payload models.CreateTransactionPayload) (*models.Transaction, error) {

	 querier, tx , err := s.repo.Begin(ctx)
	 if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx) // Ensure rollback in case of error
	var numericAmount pgtype.Numeric
    if err := numericAmount.Scan(payload.Amount); err != nil {
        return nil, fmt.Errorf("invalid amount format: %w", err)
    }
	// create the transaction
	transaction, err := querier.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		UserID: userID,
		Amount: payload.Amount,
		Type:   &payload.Type,
		Reason: &payload.Reason,
	})
	if err != nil {
		return nil, err
	}

	// Update user's total savings or withdrawals based on transaction type
	if payload.Type == "saving"{
		_, err = querier.AddToUserSavings(ctx , sqlc.AddToUserSavingsParams{
			Amount: numericAmount,
			ID: *userID,

		})
		
		
	} else if payload.Type == "withdrawal" {
			_, err = querier.AddToUserWithdrawals(ctx , sqlc.AddToUserWithdrawalsParams{
				Amount: numericAmount,
				ID: *userID,
			})
		}
	if err != nil {
			return nil, err
		}
	err = tx.Commit(ctx)
	return sqlCToAppTransaction(transaction), nil
}

func (s *Service) GetTransactions(ctx context.Context, userID int32, transactionType string) (*[]models.Transaction, error) {
	txns, err := s.repo.Do().GetTransactions(ctx, sqlc.GetTransactionsParams{
		UserID: &userID,
		TypeFilter: transactionType,
	})
	if err != nil {
		return nil, err
	}
	transactions := []models.Transaction{}
	for _, v := range txns {
		transactions = append(transactions, *sqlCToAppTransaction(v))
	}
	return &transactions, nil
}

func sqlCToAppTransaction(t sqlc.Transaction) *models.Transaction {
	// Set safe defaults
    reason := ""
    if t.Reason != nil {
        reason = *t.Reason
    }

    tType := ""
    if t.Type != nil {
        tType = *t.Type
    }
	return &models.Transaction{
		Amount:    t.Amount,
		Reason:    reason,
		Type:      tType,
		ID:        &t.ID,
		CreatedAt: t.CreatedAt.Time.String(),
	}
}

