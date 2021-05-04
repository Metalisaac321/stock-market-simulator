package postgres

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/Metalisaac321/stock-market-simulator/internal"
	"github.com/huandu/go-sqlbuilder"
)

type AcccountRepository struct {
	db *sql.DB
}

func NewAcccountRepository(db *sql.DB) *AcccountRepository {
	return &AcccountRepository{
		db: db,
	}
}

func (r *AcccountRepository) Save(ctx context.Context, account mooc.Account) error {
	acccountSQLStruct := sqlbuilder.NewStruct(new(sqlAccount)).For(sqlbuilder.PostgreSQL)
	query, args := acccountSQLStruct.InsertInto(sqlAccountTable, sqlAccount{
		Id:   account.Id().Value(),
		Cash: account.Cash().Value(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error trying to persist account on database: %v", err)
	}

	return nil
}
