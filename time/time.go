package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "gopkg.in/goracle.v2"
)

func ConnectToOeBS(conn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("goracle", conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)

	return db, nil
}

type SalaryInfo struct {
	AssignmentID int64           `db:"ASSIGNMENT_ID"`
	PayBasisID   sql.NullInt64   `db:"PAY_BASIS_ID"`
	Salary       float64         `db:"SALARY"`
	FTE          sql.NullFloat64 `db:"FTE"`
	HourRate     float64         `db:"-"`
	HourRateCoef float64         `db:"-"`
	IsHourly     bool            `db:"-"`
	DateFrom     time.Time       `db:"DATE_FROM"`
}

func main() {
	now := time.Now()
	day := 24 * time.Hour
	transferDate := now.Add(-3 * time.Hour).Unix()
	fmt.Println(now, transferDate)
	if time.Unix(transferDate, 0).Truncate(day).Equal(now.Truncate(day)) {
		fmt.Println("EQUAL DATES")
	}

	mainDB, err := ConnectToOeBS("TECH_PEPPERPOTTS/AWEFOIHASD34SDFBV9VA_@localhost:1521/OEBS14")
	if err != nil {
		log.Fatal("can't connect to oebs (main)", err)
	}
	defer mainDB.Close()

	loc, _ := time.LoadLocation("Europe/Moscow")

	now = time.Now()
	currDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	fmt.Println(now.Format(time.StampNano))

	query, args, err := sq.Select("assignment_id, NVL(salary,0) AS salary, date_from, fte,pay_basis_id").
		From("APPS.XXPER_MTS_HR_5482_ASSIGNM_PV").
		PlaceholderFormat(sq.Colon).Where("assignment_id = 126796").
		// ToSql()
		Where("effective_start_date <= ?", currDate).Where("effective_end_date >= ?", currDate).ToSql()
	if err != nil {
		fmt.Println("build query: %w", err)
	}

	fmt.Println(query, args)

	// logging.FormatLog(ctx, "", "Получение зарплаты сотрудников", query, args...)

	rows, err := mainDB.QueryxContext(context.Background(), query, args...)
	if err != nil {
		fmt.Println("query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		fmt.Println("scan")
		var info SalaryInfo
		if err = rows.StructScan(&info); err != nil {
			fmt.Println("scan: %w", err)
		}
		fmt.Println(info)
	}

}
