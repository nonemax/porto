package pq

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgre lib
	"github.com/nonemax/porto-entity"
)

// DB represents Postgres connection
type DB struct {
	DB    *sql.DB
	stmts map[int]*sql.Stmt
}

const (
	stmtGetPortID = iota
	stmtGetPort
	stmtInsertPort
	stmtUpdatePort
)

var stmtPairs = []struct {
	id   int
	stmt string
}{
	{stmtGetPortID, `SELECT ports.id
	FROM ports
		WHERE ports.unlocks=$1`},

	{stmtGetPort, `SELECT ports.name, ports.city, ports.сountry, ports.alias, ports.regions, ports.lat, ports.long, ports.province, ports.timezone, ports.code
		FROM ports 
			WHERE ports.unlocks=$1`},

	{stmtInsertPort, `INSERT INTO ports
		(name, city, сountry, alias, regions, lat, long, province, timezone, unlocks, code)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`},

	{stmtUpdatePort, `UPDATE ports SET name=$1, city=$2, сountry=$3, alias=$4, regions=$5, lat=$6, long=$7, province=$8, timezone=$9, code=$10
	WHERE id=$11`},
}

func (d *DB) prepareStmts() error {
	for _, stmtPair := range stmtPairs {
		st, err := d.DB.Prepare(stmtPair.stmt)
		if err != nil {
			return fmt.Errorf("Failed to prepare statement %s: %s", stmtPair.stmt, err.Error())
		}
		d.stmts[stmtPair.id] = st
	}

	return nil
}

// NewDB creates new Postgres connection
func NewDB(dataSource string) (*DB, error) {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	res := &DB{DB: db, stmts: map[int]*sql.Stmt{}}
	if err := res.prepareStmts(); err != nil {
		return nil, err
	}

	return res, nil
}

// SavePort insert new port or update existing
func (d *DB) SavePort(p entity.Port) error {
	var (
		lat     float32
		long    float32
		regions string
		alias   string
		id      int64
		unlocks string
	)

	if len(p.Coordinates) > 1 {
		lat = p.Coordinates[0]
		long = p.Coordinates[1]
	}
	if len(p.Regions) > 0 {
		regions = p.Regions[0]
	}
	if len(p.Alias) > 0 {
		alias = p.Alias[0]
	}
	if len(p.Unlocks) > 0 {
		unlocks = p.Unlocks[0]
	}

	err := d.stmts[stmtGetPortID].QueryRow(unlocks).Scan(&id)
	if err == sql.ErrNoRows {
		_, err = d.stmts[stmtInsertPort].Exec(
			p.Name,
			p.City,
			p.Country,
			alias,
			regions,
			lat,
			long,
			p.Province,
			p.TimeZone,
			unlocks,
			p.Code,
		)
		return nil
	} else if err != nil {
		return err
	}
	_, err = d.stmts[stmtUpdatePort].Exec(p.Name, p.City, p.Country, alias, regions, lat, long, p.Province, p.TimeZone, p.Code, id)
	if err != nil {
		return err
	}
	return nil
}

// GetPort get port by it unlock
func (d *DB) GetPort(unlock string) (entity.Port, error) {
	var (
		alias   string
		regions string
		lat     float32
		long    float32
	)
	unlocks := []string{unlock}
	port := entity.Port{Unlocks: unlocks}
	port.Alias = []string{}
	port.Regions = []string{}
	err := d.stmts[stmtGetPort].QueryRow(unlock).Scan(&port.Name, &port.City, &port.Country, &alias, &regions, &lat, &long, &port.Province, &port.TimeZone, &port.Code)
	if err == sql.ErrNoRows {
		return port, err
	} else if err != nil {
		return port, err
	}
	if len(alias) > 0 {
		port.Alias = []string{alias}
	}
	if len(regions) > 0 {
		port.Regions = []string{regions}
	}
	port.Coordinates = []float32{lat, long}
	return port, nil
}
