package platform

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Metalisaac321/stock-market-simulator/internal/account"
	accountMap "github.com/Metalisaac321/stock-market-simulator/internal/account/mapper"
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
		return fmt.Errorf("error trying to persist account on database: %v", err)
	}

	return nil
}

func (acccountRepository *AcccountRepository) SearchAll(ctx context.Context) ([]account.Account, error) {
	query := sqlbuilder.NewSelectBuilder().Select("*").From(sqlAccountTable).String()
	rows, err := acccountRepository.db.QueryContext(ctx, query)

	if err != nil {
		return []account.Account{}, fmt.Errorf("error trying to query account on database: %v", err)
	}

	var accounts []account.Account

	defer rows.Close()
	for rows.Next() {
		accountDto := accountMap.AccountDto{}
		err = rows.Scan(&accountDto.Id, &accountDto.Cash)
		if err != nil {
			fmt.Println(err)
			return []account.Account{}, fmt.Errorf("error trying to query account on database: %v", err)
		}

		accounts = append(accounts, accountMap.ToDomain(accountDto))
	}

	return accounts, nil
}
