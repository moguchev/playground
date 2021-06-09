package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type TimeZoneOffset struct {
	offset int8
}

func (t *TimeZoneOffset) DecodeText(src string) error {
	const (
		utc           = "UTC"
		pgxTimeLayout = "03:00:00"
		maxTimezone   = 14
		minTimezone   = -12
	)

	var (
		parse int64
		err   error
		tmp   time.Time
	)
	if strings.HasPrefix(src, utc) {
		parse, err = strconv.ParseInt(strings.TrimPrefix(src, utc), 10, 8)
	} else {
		tmp, err = time.Parse(pgxTimeLayout, src)
		parse = int64(tmp.Hour())
	}
	if err != nil {
		return err
	}

	if parse >= minTimezone && parse <= maxTimezone {
		t.offset = int8(parse)
		return nil
	}
	return fmt.Errorf("wrong timezone: %s", src)
}

func (t *TimeZoneOffset) DecodeTime(src time.Time) error {
	t.offset = int8(src.Hour())
	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *TimeZoneOffset) Scan(src interface{}) error {
	if src == nil {
		dst = nil
		return nil
	}

	switch src := src.(type) {
	case string:
		return dst.DecodeText(src)
	case time.Time:
		return dst.DecodeTime(src)
	default:
		return fmt.Errorf("cannot scan %T", src)
	}
}

type Tag struct {
	ID    int64
	Name  string
	Alias string
}

type Tags []Tag

func (t Tags) Equal(o Tags) bool {
	if len(t) != len(o) {
		return false
	}

	m := make(map[int64]bool, len(o))
	for i := range o {
		m[o[i].ID] = true
	}

	for i := range t {
		if ok := m[t[i].ID]; !ok {
			return false
		}
	}

	return true
}

// Value implements the database/sql/driver Valuer interface.
func (src Tag) Value() (driver.Value, error) {
	const api = "Tag.Value"

	return src.ID, nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Tag) Scan(src interface{}) error {
	const api = "Tag.Scan"
	fmt.Printf(api+": %T", src)
	if src == nil {
		dst = nil
		return nil
	}
	var (
		id  int64
		err error
	)
	switch src.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		id = src.(int64)
	case []byte, string:
		id, err = strconv.ParseInt(src.(string), 10, 64)
	default:
		err = fmt.Errorf("cannot scan %T", src)
	}

	if err != nil {
		return errors.Wrap(err, api)
	}

	*dst = Tag{ID: id}
	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Tags) Value() (driver.Value, error) {
	const api = "Tags.Value"
	if src == nil {
		return nil, nil
	}

	ids := make([]int64, 0, len(src))
	for i := range src {
		ids = append(ids, src[i].ID)
	}
	// optional
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	v, err := pq.Array(ids).Value()
	return v, errors.Wrap(err, api)
}

// Scan implements the database/sql Scanner interface.
func (dst *Tags) Scan(src interface{}) error {
	const api = "Tags.Scan"
	log.Printf("%#v, %T", src, src)
	if src == nil {
		return nil
	}

	var ids []int64
	if err := pq.Array(&ids).Scan(src); err != nil {
		return errors.Wrap(err, api)
	}

	*dst = make(Tags, len(ids))
	for i := range ids {
		(*dst)[i] = Tag{ID: ids[i]}
	}

	return nil
}

func main() {
	var (
	// t, tp time.Time
	// s, sp string
	// m, mp TimeZoneOffset
	)
	// ctx := context.Background()
	conn := "host=localhost port=6432 user=datagateway-user password=QuePhoh7xeing7U dbname=datagateway sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("ping: %w", err)
	}

	// if err := db.QueryRow("SELECT time_zone_offset FROM deliveryvariant WHERE id = $1", 20703543498000).Scan(&t); err != nil {
	// 	log.Printf("postgres: time: %s", err)
	// }
	// if err := db.QueryRow("SELECT time_zone_offset FROM deliveryvariant WHERE id = $1", 20703543498000).Scan(&s); err != nil {
	// 	log.Printf("postgres: string: %s", err)
	// }
	// if err := db.QueryRow("SELECT time_zone_offset FROM deliveryvariant WHERE id = $1", 20703543498000).Scan(&m); err != nil {
	// 	log.Printf("postgres: string: %s", err)
	// }

	///////////////////////////////////////////////
	pgxdb, err := pgx.Connect(pgx.ConnConfig{
		Host:     "0.0.0.0",
		Port:     6432,
		Database: "datagateway",
		User:     "datagateway-user",
		Password: "QuePhoh7xeing7U",
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pgxdb.Close()
	// _, err = pgxdb.Prepare("get dv by id 1", "SELECT time_zone_offset FROM deliveryvariant WHERE id = $1")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := pgxdb.QueryRow("SELECT time_zone_offset FROM deliveryvariant WHERE id = $1", 20703543498000).Scan(&sp); err != nil {
	// 	log.Printf("pgx: string: %s", err)
	// }
	// if err := pgxdb.QueryRow("SELECT time_zone_offset FROM deliveryvariant WHERE id = $1", 20703543498000).Scan(&tp); err != nil {
	// 	log.Printf("pgx: time: %s", err)
	// }
	var tags Tags
	if err := pgxdb.QueryRow("SELECT tags FROM deliveryvariant WHERE id = $1", 1011000000002420).Scan(&tags); err != nil {
		log.Printf("pgx: err: %s", err)
	}

	// var zone TimeZoneOffset
	// zone.DecodeText("UTC+3")
	log.Println(tags)

}
