package users

import (
	"encore.dev/storage/sqldb"
	"github.com/omid9h/encore.app.bp/users/repo"
)

var _ = sqldb.NewDatabase("user", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

var usersdb = sqldb.Named("users").Stdlib()

//encore:service
type Service struct {
	repo *repo.Queries
	// Issuer defines the application name for creating token
	Issuer string
}

func NewService(repo *repo.Queries) *Service {
	return &Service{
		repo:   repo,
		Issuer: "encore.app.bp",
	}
}

// initService initializes the user service.
// It is automatically called by Encore on service startup.
func initService() (*Service, error) {
	return NewService(repo.New(usersdb)), nil
}
