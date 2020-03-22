package config

import (
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/pkg/errors"
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

type AsclepiusDataStore struct {
	Endpoint string
	User string
	Password string
	DatabaseName string
}

type GlobalConfig struct {
	Messenger MessengerConfig
	Watson WatsonConfig
	Twilio TwilioConfig
	Storage AsclepiusDataStore
}


func getGlobalConfigFromVault() (*GlobalConfig, error) {
	vaultAddress := os.Getenv("VAULT_ADDRESS")
	vaultToken := os.Getenv("VAULT_TOKEN")

	client, err := api.NewClient(&api.Config{Address: vaultAddress})
	if err != nil {
		return nil, errors.Wrap(err, "error at create client for vault api")
	}

	client.SetToken(vaultToken)

	tape := client.Logical()

	asclepius := "kv/data/asclepius"

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

	database, err := tape.Read(asclepius + "/database")
	if err != nil {
		return nil, errors.Wrap(err, "cannot read kv twilio")
	}

	config := new(GlobalConfig)

	if messenger.Data["data"] != nil {
		data := messenger.Data["data"].(map[string]interface{})
		config.Messenger = MessengerConfig{
			AccessToken:  data["accessToken"].(string),
			VerifySecret: data["verifySecret"].(string),
			PageID:       data["pageID"].(string),
		}
	}

	if watson.Data["data"] != nil {
		data := watson.Data["data"].(map[string]interface{})
		config.Watson = WatsonConfig{
			URL:         data["url"].(string),
			Version:     data["version"].(string),
			AssistantID: data["assistantID"].(string),
			APIKey:      data["apiKey"].(string),
		}
	}

	if twilio.Data["data"] != nil {
		data := twilio.Data["data"].(map[string]interface{})
		config.Twilio = TwilioConfig{
			AccountID: data["accountID"].(string),
			AuthToken: data["authToken"].(string),
		}
	}
	
	if database.Data["data"] != nil {
		data := database.Data["data"].(map[string]interface{})
		config.Storage = AsclepiusDataStore{
			Endpoint:     data["endpoint"].(string),
			User:         data["user"].(string),
			Password:     data["password"].(string),
			DatabaseName: data["databaseName"].(string),
		}
	}

	return config, nil
}

func LoadFromVault() (*GlobalConfig, error) {
	return getGlobalConfigFromVault()
}
