package user

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/rhperera/go-base/domain"
)

type mssqlRepo struct {
	db *sql.DB
}

func (m *mssqlRepo) GetByID(id int64) (*domain.User, error) {
	panic("implement me")
}

func NewMSSqlRepo(d *sql.DB) domain.UserRepo  {
	return &mssqlRepo{db: d}
}