package pg

import (
	"database/sql"
	"fmt"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const containersTableName = "containers"

func NewContainerQ(db *pgdb.DB) data.ContainerQ {
	return &ContainerQ{
		db:  db.Clone(),
		sql: sq.Select("b.*").From(fmt.Sprintf("%s as b", containersTableName)),
	}
}

type ContainerQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *ContainerQ) New() data.ContainerQ {
	return NewContainerQ(q.db)
}

func (q *ContainerQ) Get() (*data.Container, error) {
	var result data.Container
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *ContainerQ) Select() ([]data.Container, error) {
	var result []data.Container
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *ContainerQ) Insert(value data.Container) error {
	clauses := structs.Map(value)

	stmt := sq.Insert(containersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Exec(stmt)

	return err
}
func (q *ContainerQ) Page(pageParams pgdb.OffsetPageParams) data.ContainerQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *ContainerQ) FilterByAddress(ids ...string) data.ContainerQ {
	q.sql = q.sql.Where(sq.Eq{"b.owner_address": ids})
	return q
}
func (q *ContainerQ) DelById(ids ...int64) error {
	s := sq.Delete(containersTableName).Where(sq.Eq{"id": ids})
	err := q.db.Exec(s)
	return err
}
func (q *ContainerQ) FilterByID(ids ...int64) data.ContainerQ {
	q.sql = q.sql.Where(sq.Eq{"b.id": ids})
	return q
}
