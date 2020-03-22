package main

import (
	"github.com/minskylab/asclepius"
	"github.com/minskylab/asclepius/bot"
	"github.com/minskylab/asclepius/config"
	"github.com/pkg/errors"
)

func main() {
	conf, err := config.LoadFromVault()
	if err != nil {
		panic(errors.Cause(err))
	}

	core, err := asclepius.NewCore()
	if err != nil {
		panic(errors.Cause(err))
	}

	jarvis, err := bot.NewBot(conf, core)
	if err != nil {
		panic(errors.Cause(err))
	}

	if err := jarvis.Run(); err != nil {
		panic(errors.Cause(err))
	}


}