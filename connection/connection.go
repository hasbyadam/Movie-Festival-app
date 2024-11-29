package connection

import (
	"database/sql"
	"fmt"
	"movie-festival-app/entity"

	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"
)

// Postgres
func Postgres(opts entity.Pg) *sql.DB {
	var pemFile string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s sslmode=%s",
		opts.Host, opts.Port, opts.User, opts.Password, opts.Dbname, pemFile, opts.Sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		zap.S().DPanic(err)
	}

	err = db.Ping()
	if err != nil {
		zap.S().DPanic(err)
	}

	// Setting database connection config
	db.SetMaxOpenConns(opts.MaxOpenConnection)
	db.SetMaxIdleConns(opts.MaxIdleConnection)
	db.SetConnMaxLifetime(opts.MaxConnectionLifetime)

	zap.S().Info("Connected to PG DB Server: ", opts.Host, " at port:", opts.Port, " successfully!")

	return db
}
