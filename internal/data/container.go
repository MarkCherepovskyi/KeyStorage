package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

type ContainerQ interface {
	New() ContainerQ

	Get() (*Container, error)
	Select() ([]Container, error)

	Insert(data Container) (int64, error)
	Page(pageParams pgdb.OffsetPageParams) ContainerQ
	DelById(id ...int64) error
	FilterByAddress(id ...string) ContainerQ
	FilterByID(id ...int64) ContainerQ
}

type Container struct {
	ID        int64    `db:"id" structs:"-"`
	Owner     string   `db:"owner" structs:"owner"`
	Recipient []string `db:"recipient" structs:"recipient"`
	Group     bool     `db:"group" structs:"group"`
	Tag       string   `db:"tag" structs:"tag"`
	Container string   `db:"container" structs:"container"`
}
