package textbots_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_bots_execute"
)

func init() {
	textbots_botsCmd.AddCommand(textbots_bots_execute.Cmdtextbots_bots_execute())
	textbots_botsCmd.Short = utils.GenerateCustomDescription(textbots_botsCmd.Short, textbots_bots_execute.Description, )
	textbots_botsCmd.Long = textbots_botsCmd.Short
}
