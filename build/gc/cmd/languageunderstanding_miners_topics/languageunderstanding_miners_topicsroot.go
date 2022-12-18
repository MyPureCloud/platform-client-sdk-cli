package languageunderstanding_miners_topics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding_miners_topics_phrases"
)

func init() {
	languageunderstanding_miners_topicsCmd.AddCommand(languageunderstanding_miners_topics_phrases.Cmdlanguageunderstanding_miners_topics_phrases())
	languageunderstanding_miners_topicsCmd.Short = utils.GenerateCustomDescription(languageunderstanding_miners_topicsCmd.Short, languageunderstanding_miners_topics_phrases.Description, )
	languageunderstanding_miners_topicsCmd.Long = languageunderstanding_miners_topicsCmd.Short
}
