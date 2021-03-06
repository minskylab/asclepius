package asclepius

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect"
	"github.com/minskylab/asclepius/config"
	"github.com/minskylab/asclepius/ent"
	"github.com/minskylab/asclepius/ent/migrate"
	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

type Core struct {
	client *ent.Client
}

func NewCore(config *config.GlobalConfig) (*Core, error) {
	core := new(Core)
	if err := core.openClient(config); err != nil {
		return nil, err
	}
	return core, nil
}

func (core *Core) openClient(config *config.GlobalConfig) error {
	var err error

	user := config.Storage.User
	pass := config.Storage.Password
	endpoint := config.Storage.Endpoint
	database := config.Storage.DatabaseName

	// e.g. "root:pass@tcp(localhost:3306)/test"
	uri :=fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, endpoint, database)

	core.client, err = ent.Open(dialect.MySQL, uri )
	if err != nil {
		return errors.Wrap(err, "failed opening connection to mysql database")
	}

	if err = core.client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		return errors.Wrap(err, "error at create schema")
	}
	return nil
}

func (core *Core) done() error {
	return core.client.Close()
}

func (core *Core) Done() error {
	return core.done()
}