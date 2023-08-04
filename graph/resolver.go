package graph

import (
	"database/sql"

	"github.com/andersjbe/quin/internal/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *database.Queries
	Conn *sql.DB
}
