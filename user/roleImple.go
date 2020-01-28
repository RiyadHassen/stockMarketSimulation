package user

import sql2 "database/sql"

type RoleRepository struct {
	sql *sql2.DB
}

func NewRoleRepository(conn *sql2.DB) *RoleRepository {
	return &RoleRepository{sql:conn}
} 