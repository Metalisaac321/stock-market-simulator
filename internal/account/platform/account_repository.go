package platform

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Metalisaac321/stock-market-simulator/internal/account"
	"github.com/huandu/go-sqlbuilder"
)

const (
	sqlAccountTable = "account"
)

type sqlAccount struct {
	Id   string `db:"id"`
	Cash uint   `db:"cash"`
}
type AcccountRepository struct {
	db *sql.DB
}

func NewAcccountRepository(db *sql.DB) *AcccountRepository {
	return &AcccountRepository{
		db: db,
	}
}

func (acccountRepository *AcccountRepository) Save(ctx context.Context, account account.Account) error {
	acccountSQLStruct := sqlbuilder.NewStruct(new(sqlAccount)).For(sqlbuilder.PostgreSQL)
	query, args := acccountSQLStruct.InsertInto(sqlAccountTable, sqlAccount{
		Id:   account.Id().Value(),
		Cash: account.Cash().Value(),
	}).Build()

	_, err := acccountRepository.db.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error trying to persist account on database: %v", err)
	}

	return nil
}

func (acccountRepository *AcccountRepository) SearchAll(ctx context.Context) ([]account.Account, error) {
	a, _ := account.NewAccount("de82924b-8e16-431b-b945-5e4eaf6ecc29", 100)

	accounts := []account.Account{a}

	return accounts, nil
}
