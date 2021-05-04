package postgres

const (
	sqlAccountTable = "account"
)

type sqlAccount struct {
	Id   string `db:"id"`
	Cash uint   `db:"cash"`
}
