package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"text/template"
	"time"

	_ "github.com/lib/pq"
	"github.com/pooulad/go-http-server/pkg/config"
)

const connString = "postgres://{{.Username}}:{{.Password}}@{{.Host}}:{{.Port}}/{{.Database}}?sslmode=disable"

func buildCnnectionStringOrPanic(cnf config.Postgres) string {
	sb := strings.Builder{}
	temp := template.Must(template.New("ConnString").Parse(connString))
	err := temp.Execute(&sb, cnf)
	if err != nil {
		panic(err)
	}
	return sb.String()
}

func NewPostgres(cnf config.Postgres) (*sql.DB, error) {
	conn := buildCnnectionStringOrPanic(cnf)

	//" pg driver : github.com/lib/pq"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	errDB := db.Ping()
	if errDB != nil {
		fmt.Print("Faild to ping the database")
		return db, fmt.Errorf("could not ping database %w", errDB)
	}

	db.SetConnMaxLifetime(time.Second)
	db.SetConnMaxIdleTime(30 * time.Second)
	return db, nil
}
