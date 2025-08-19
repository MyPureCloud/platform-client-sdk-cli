package integrations_speech_lexv2

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lexv2_bot"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_lexv2_bots"
)

func init() {
	integrations_speech_lexv2Cmd.AddCommand(integrations_speech_lexv2_bot.Cmdintegrations_speech_lexv2_bot())
	integrations_speech_lexv2Cmd.AddCommand(integrations_speech_lexv2_bots.Cmdintegrations_speech_lexv2_bots())
	integrations_speech_lexv2Cmd.Short = utils.GenerateCustomDescription(integrations_speech_lexv2Cmd.Short, integrations_speech_lexv2_bot.Description, integrations_speech_lexv2_bots.Description, )
	integrations_speech_lexv2Cmd.Long = integrations_speech_lexv2Cmd.Short
}
