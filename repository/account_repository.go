package repository

import (
	"database/sql"
	"fmt"

	"github.com/adycahyoputro/manage-bank/model/dto"
	"github.com/adycahyoputro/manage-bank/model/entity"
)

type AccountRepository interface {
	FindAccountByUserID(userID string) (*entity.Account, error)
	FindAccountByAccountID(accountID string) (*entity.Account, error)
	CreateAccount(newAccount *entity.Account, tx *sql.Tx) (*entity.Account, error)
	UpdateAccount(updateAccount *dto.AccountRequest, tx *sql.Tx) (*dto.AccountRequest, error)
	UpdateIsActiveAccount(updateIsActive *dto.LogoutRequest) error
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (repo *accountRepository) FindAccountByUserID(userID string) (*entity.Account, error) {
	var account entity.Account
	stmt, err := repo.db.Prepare("SELECT id, user_id, owner, balance, currency, created_at FROM accounts WHERE user_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&account.ID, &account.UserID, &account.Owner, &account.Balance, &account.Currency, &account.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account with user id %s not found", userID)
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repo *accountRepository) FindAccountByAccountID(accountID string) (*entity.Account, error) {
	var account entity.Account
	stmt, err := repo.db.Prepare("SELECT id, user_id, owner, balance, currency, created_at, is_active FROM accounts WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(accountID).Scan(&account.ID, &account.UserID, &account.Owner, &account.Balance, &account.Currency, &account.CreatedAt, &account.IsActive)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account with account id %s not found", accountID)
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repo *accountRepository) CreateAccount(newAccount *entity.Account, tx *sql.Tx) (*entity.Account, error) {
	stmt, err := repo.db.Prepare("INSERT INTO accounts (id, user_id, owner, balance, currency, created_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to create account : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newAccount.ID, newAccount.UserID, newAccount.Owner, newAccount.Balance, newAccount.Currency, newAccount.CreatedAt).Scan(&newAccount.ID)

	validate(err, "create account", tx)

	return newAccount, nil
}

func (repo *accountRepository) UpdateAccount(updateAccount *dto.AccountRequest, tx *sql.Tx) (*dto.AccountRequest, error) {
	stmt, err := repo.db.Prepare("UPDATE accounts SET balance = $1 WHERE id = $2")
	if err != nil {
		return nil, fmt.Errorf("failed to update account : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateAccount.Balance, updateAccount.AccountID)

	validate(err, "update account", tx)

	return updateAccount, nil
}

func (repo *accountRepository) UpdateIsActiveAccount(updateIsActive *dto.LogoutRequest) error {
	stmt, err := repo.db.Prepare("UPDATE accounts SET is_active = $1 WHERE user_id = $2")
	if err != nil {
		return fmt.Errorf("failed to update account : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateIsActive.IsActive, updateIsActive.UserID)
	if err != nil {
		return fmt.Errorf("failed to update account : %w", err)
	}

	return nil
}
