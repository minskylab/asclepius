package asclepius

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/minskylab/asclepius/ent"
	"github.com/pkg/errors"
	// log "github.com/sirupsen/logrus"
)

type Core struct {
	client *ent.Client
}

func NewCore() (*Core, error) {
	core := new(Core)
	if err := core.openClient(); err != nil {
		return nil, err
	}
	return core, nil
}

func (core *Core) openClient() error {
	var err error
	core.client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return errors.Wrap(err, "failed opening connection to sqlite")
	}
	return nil
}

func (core *Core) done() error {
	return core.client.Close()
}