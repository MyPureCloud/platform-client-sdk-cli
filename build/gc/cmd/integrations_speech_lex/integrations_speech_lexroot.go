package integrations_speech_lex

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lex_bot"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lex_bots"
)

func init() {
	integrations_speech_lexCmd.AddCommand(integrations_speech_lex_bot.Cmdintegrations_speech_lex_bot())
	integrations_speech_lexCmd.AddCommand(integrations_speech_lex_bots.Cmdintegrations_speech_lex_bots())
	integrations_speech_lexCmd.Short = utils.GenerateCustomDescription(integrations_speech_lexCmd.Short, integrations_speech_lex_bot.Description, integrations_speech_lex_bots.Description, )
	integrations_speech_lexCmd.Long = integrations_speech_lexCmd.Short
}
