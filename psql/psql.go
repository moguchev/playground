package main

import (
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes"

	"github.com/jmoiron/sqlx"

	"database/sql"
	"database/sql/driver"

	"github.com/jackc/pgx/v4/stdlib"
	"github.com/opencensus-integrations/ocsql"
)

type Config struct {
	Connection      string        `yaml:"postgresql"`
	MaxOpenConn     int           `yaml:"max_open_conn"`
	MaxIdleConn     int           `yaml:"max_idle_conn"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`
	opts            options
}

type options struct {
	Wrapper func(driver.Connector) driver.Connector
}

func TelemetryWrapper(drv driver.Connector) driver.Connector {
	return ocsql.WrapConnector(drv, ocsql.WithOptions(ocsql.TraceOptions{
		AllowRoot:    true,
		Ping:         true,
		RowsClose:    true,
		RowsAffected: true,
		LastInsertID: true,
		Query:        true,
	}))
}

func WithWrapper(cfg Config, wrapper func(drv driver.Connector) driver.Connector) Config {
	if cfg.opts.Wrapper == nil {
		cfg.opts.Wrapper = wrapper
	} else {
		cfg.opts.Wrapper = func(drv driver.Connector) driver.Connector {
			return wrapper(cfg.opts.Wrapper(drv))
		}
	}
	return cfg
}

var DefaultWrapper = TelemetryWrapper

func (cfg Config) CreateDB() (*sqlx.DB, error) {
	var (
		ctor driver.Connector
	)

	drv := stdlib.GetDefaultDriver().(*stdlib.Driver)

	ctor, err := drv.OpenConnector(cfg.Connection)
	if err != nil {
		return nil, err
	}

	if cfg.opts.Wrapper == nil {
		cfg.opts.Wrapper = DefaultWrapper
	}

	ctor = cfg.opts.Wrapper(ctor)

	db := sql.OpenDB(ctor)

	if cfg.MaxConnLifetime != 0 {
		db.SetConnMaxLifetime(cfg.MaxConnLifetime)
	}

	if cfg.MaxIdleConn != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConn)
	}

	if cfg.MaxOpenConn != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	return sqlx.NewDb(db, "pgx"), nil
}

type SalaryChangeInfo struct {
	AssignmentID int64     `db:"assignment_id"`
	Date         time.Time `db:"last_change_date"`
	Diff         float64   `db:"sal_diff"`
}

var getSalaryChangeQuery = sq.Select("assignment_id, last_change_date, sal_diff").
	From("salarychanges").
	PlaceholderFormat(sq.Dollar)

func main() {
	conn := "host=localhost port=5432 user=revision password=gd1dd2jkr87ds dbname=oebs sslmode=disable"

	cfg := Config{
		Connection: conn,
	}

	db, err := cfg.CreateDB()
	if err != nil {
		log.Fatal("crete db: %w", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("ping: %w", err)
	}

	assignments := []int64{355716}
	orSt := make(sq.Or, 0, len(assignments))
	for k := range assignments {
		orSt = append(orSt, sq.Eq{"assignment_id": fmt.Sprint(assignments[k])})
	}

	query, args, err := getSalaryChangeQuery.Where(orSt).ToSql()
	if err != nil {
		log.Fatal("build query: %w", err)
	}

	rows, err := db.Queryx(query, args...)
	if err != nil {
		log.Fatal("query: %w", err)
	}

	defer rows.Close()
	var info SalaryChangeInfo
	for rows.Next() {
		if err = rows.StructScan(&info); err != nil {
			log.Fatal("scan: %w", err)
		}
		log.Println(info)
	}

	dateFromProto, err := ptypes.TimestampProto(info.Date.UTC())
	if err != nil {
		log.Fatal("TimestampProto: %w", err)
	}
	log.Println(dateFromProto)

	dateFrom, err := ptypes.Timestamp(dateFromProto)
	if err != nil {
		log.Fatal("Timestamp: %w", err)
	}

	log.Println(dateFrom.Format("02.01.2006"))
}
