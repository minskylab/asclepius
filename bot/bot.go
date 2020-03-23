package bot

import (
	"github.com/minskylab/asclepius"
	"github.com/minskylab/asclepius/config"
	"github.com/minskylab/asclepius/emitter"
	neo "github.com/minskylab/neocortex"

	"github.com/pkg/errors"
)

type Bot struct {
	c *config.GlobalConfig
	engine *neo.Engine
	emitter *emitter.Emitter
	core *asclepius.Core
}

func NewBot(c *config.GlobalConfig, core *asclepius.Core) (*Bot, error) {
	bot := new(Bot)
	var err error

	bot.c = c
	bot.core = core

	brain, err := loadCognitive(c)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to load the cognitive service")
	}

	channels, err := loadChannels(c)
	if err != nil {
		return nil, errors.Wrap(err, "error at load channels")
	}

	bot.engine, err = neo.Default(nil, brain, channels...)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create neocortex engine default")
	}

	bot.emitter, err = emitter.NewEmitter(c.Twilio.AccountID, c.Twilio.AuthToken)
	if err != nil {
		return nil, errors.Wrap(err, "error at config emitter, currently only twilio")
	}

	bot.actionsForFb(channels[0])

	return bot, nil
}

func (bot *Bot) Run() error {
	if err := bot.engine.Run(); err != nil {
		return errors.Wrap(err, "error at run bot engine")
	}
	return nil
}

func (bot *Bot) Done() error {
	return bot.core.Done()
}