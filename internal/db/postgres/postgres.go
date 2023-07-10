package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"text/template"

	"github.com/samanxzrpor/http-server/pkg/config"
)

const ConnString = "posgtres://{{.User}}:{{.Pass}}@{{.Host}}:{{.Port}}/{{.Db}}@sslmode=diable"

func BuildConnStringOrPanic(cnf config.Database) string {
	cb := strings.Builder{}
	temp := template.Must(template.New("Conn").Parse(ConnString))
	if err := temp.Execute(&cb, cnf); err != nil {
		panic(err)
	}
	return cb.String()
}

func NewPostgres(cnf config.Database) (*sql.DB, error) {

	conn := BuildConnStringOrPanic(cnf)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Print("Failed to Ping Database")
		return db, fmt.Errorf("Could not ping to database %w", err)
	}

	return db, nil
}
