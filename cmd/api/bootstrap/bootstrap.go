package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/Metalisaac321/stock-market-simulator/internal/account/application"
	"github.com/Metalisaac321/stock-market-simulator/internal/account/platform"
	"github.com/Metalisaac321/stock-market-simulator/internal/platform/server"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 8080
	dbUser = "postgres"
	dbPass = "SuperSecretPassword"
	dbHost = "localhost"
	dbPort = "5433"
	dbName = "db"
)

func Run() error {
	//postgresql://postgres:SuperSecretPassword@db:5432/db'
	db, err := sql.Open("postgres", "postgres://postgres:SuperSecretPassword@localhost:5433/db?sslmode=disable")
	if err != nil {
		fmt.Println("error de conexión")
		return err
	}

	postgresAccountRepository := platform.NewAcccountRepository(db)
	createAccount := application.NewCreateAccount(postgresAccountRepository)
	searchAllAccounts := application.NewSearchAllAccounts(postgresAccountRepository)

	srv := server.New(host, port, createAccount, searchAllAccounts)
	return srv.Run()
}
