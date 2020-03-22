package main

import (
	"github.com/minskylab/asclepius/config"
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/channels/facebook"
	"github.com/pkg/errors"
)

func loadChannels(config *config.GlobalConfig) ([]neo.CommunicationChannel, error) {
	channels := make([]neo.CommunicationChannel, 0)

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: config.Messenger.AccessToken,
		VerifyToken: config.Messenger.VerifySecret,
		PageID:      config.Messenger.PageID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error at create new facebook channel")
	}

	channels = append(channels, fb)

	return channels, nil
}
