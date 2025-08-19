package integrations_speech_lexv2_bot

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lexv2_bot_alias"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lexv2_bot_aliases"
)

func init() {
	integrations_speech_lexv2_botCmd.AddCommand(integrations_speech_lexv2_bot_alias.Cmdintegrations_speech_lexv2_bot_alias())
	integrations_speech_lexv2_botCmd.AddCommand(integrations_speech_lexv2_bot_aliases.Cmdintegrations_speech_lexv2_bot_aliases())
	integrations_speech_lexv2_botCmd.Short = utils.GenerateCustomDescription(integrations_speech_lexv2_botCmd.Short, integrations_speech_lexv2_bot_alias.Description, integrations_speech_lexv2_bot_aliases.Description, )
	integrations_speech_lexv2_botCmd.Long = integrations_speech_lexv2_botCmd.Short
}
