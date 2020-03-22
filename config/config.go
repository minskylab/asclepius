package config

import (
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
	// log "github.com/sirupsen/logrus"
)

type MessengerConfig struct {
	AccessToken string
	VerifySecret string
	PageID string
}

type WatsonConfig struct {
	URL string
	Version string
	AssistantID string
	APIKey string
}

type TwilioConfig struct {
	AccountID string
	AuthToken string
}

type GlobalConfig struct {
	Messenger MessengerConfig
	Watson WatsonConfig
	Twilio TwilioConfig
}


func getGlobalConfig() (*GlobalConfig, error) {
	vaultAddress := os.Getenv("VAULT_ADDRESS")
	vaultToken := os.Getenv("VAULT_TOKEN")

	client, err := api.NewClient(&api.Config{Address: vaultAddress})
	if err != nil {
		return nil, errors.Wrap(err, "error at create client for vault api")
	}

	client.SetToken(vaultToken)

	tape := client.Logical()

	asclepius := "secrets/kv/asclepius"

	messenger, err := tape.Read(asclepius + "/facebook/messenger")
	if err != nil {
		return nil, errors.Wrap(err, "cannot read kv facebook messenger")
	}

	watson, err := tape.Read(asclepius + "/watson")
	if err != nil {
		return nil, errors.Wrap(err, "cannot read kv watson")
	}

	twilio, err := tape.Read(asclepius + "/twilio")
	if err != nil {
		return nil, errors.Wrap(err, "cannot read kv twilio")
	}

	config := new(GlobalConfig)

	config.Messenger.AccessToken = messenger.Data["accessToken"].(string)
	config.Messenger.PageID = messenger.Data["pageID"].(string)
	config.Messenger.VerifySecret = messenger.Data["verifySecret"].(string)

	config.Watson.APIKey = watson.Data["apiKey"].(string)
	config.Watson.AssistantID = watson.Data["assistantID"].(string)
	config.Watson.URL = watson.Data["url"].(string)
	config.Watson.Version = watson.Data["version"].(string)

	config.Twilio.AccountID = twilio.Data["accountID"].(string)
	config.Twilio.AuthToken = twilio.Data["authToken"].(string)

	return config, nil
}

func LoadFromVault() (*GlobalConfig, error) {
	return getGlobalConfig()
}
