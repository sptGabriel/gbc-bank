package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sptGabriel/banking/app"
	"github.com/sptGabriel/banking/app/domain/entities"
	"github.com/sptGabriel/banking/app/domain/repositories"
	"github.com/sptGabriel/banking/app/domain/vos"
)

type accountRepository struct {
	conn *pgxpool.Pool
}

func NewAccountRepository(c *pgxpool.Pool) repositories.AccountRepository {
	return &accountRepository{c}
}

func (r accountRepository) DoesAccountExistByCPF(ctx context.Context, cpf vos.CPF) (bool, error) {
	var query = `SELECT EXISTS(SELECT id FROM accounts WHERE cpf = $1)`
	accountExists := false
	if err := r.conn.QueryRow(ctx, query, cpf.String()).Scan(&accountExists); err != nil {
		return false, err
	}
	return accountExists, nil
}

func (r accountRepository) Create(ctx context.Context, account *entities.Account) error {
	var query = `
		INSERT INTO
			accounts (id, name, cpf, secret, balance)
		VALUES ($1, $2, $3, $4, $5)
		`
	if _, err := r.conn.Exec(ctx, query,
		account.Id.String(),
		account.Name.String(),
		account.CPF.String(),
		account.Secret.String(),
		account.Balance,
	); err != nil {
		return err
	}

	return nil
}

func (r accountRepository) UpdateAccountBalance(ctx context.Context, accountId *vos.AccountId) error {
	panic("implement me")
}

func (r accountRepository) GetByCPF(ctx context.Context, cpf vos.CPF) (*entities.Account, error) {
	var query = "SELECT id, name, cpf, secret, balance, created_at FROM accounts WHERE cpf = $1"

	var account entities.Account

	if err := r.conn.QueryRow(ctx, query, cpf).Scan(
		&account.Id,
		&account.Name,
		&account.CPF,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, app.NewNotFoundError("account")
		}
		return nil, err
	}

	return &account, nil
}

func (r accountRepository) GetAccountBalance(ctx context.Context, accId vos.AccountId) (*entities.Account, error) {
	var query = "SELECT balance FROM accounts WHERE id = $1"

	var account entities.Account

	if err := r.conn.QueryRow(ctx, query, accId.String()).Scan(
		&account.Id,
		&account.Name,
		&account.CPF,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, app.NewNotFoundError("account")
		}
		return nil, err
	}

	return &account, nil
}

func (r accountRepository) GetAll(ctx context.Context) ([]entities.Account, error) {
	var query = `select id, name, cpf, created_at from accounts`

	var accounts []entities.Account

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return accounts, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account entities.Account
		err := rows.Scan(
			&account.Id,
			&account.Name,
			&account.CPF,
			&account.Secret,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
