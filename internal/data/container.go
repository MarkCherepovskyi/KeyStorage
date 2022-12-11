package data

import (
	"github.com/lib/pq"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type ContainerQ interface {
	New() ContainerQ

	Get() (*Container, error)
	Select() ([]Container, error)

	Insert(data Container) error
	Page(pageParams pgdb.OffsetPageParams) ContainerQ
	DelById(id ...int64) error
	FilterByAddress(id ...string) ContainerQ
	FilterByID(id ...int64) ContainerQ
}

type Container struct {
	ID        int64          `db:"id" structs:"-"`
	Owner     string         `db:"owner" structs:"owner"`
	Recipient pq.StringArray `db:"recipient" structs:"recipient"`
	//Group     bool     `db:"group" structs:"group"`
	Tag       string `db:"tag" structs:"tag"`
	Container []byte `db:"container" structs:"container"`
}
