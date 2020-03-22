package bot

import (
	"github.com/minskylab/asclepius/config"
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/cognitive/watson"
)

func loadCognitive(config *config.GlobalConfig) (neo.CognitiveService, error) {
	return watson.NewCognitive(watson.NewCognitiveParams{
		Url:         config.Watson.URL,
		Version:     config.Watson.Version,
		AssistantID: config.Watson.AssistantID,
		ApiKey:      config.Watson.APIKey,
	})
}
