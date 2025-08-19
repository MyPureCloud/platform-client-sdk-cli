package textbots_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_bots_execute"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_bots_search"
)

func init() {
	textbots_botsCmd.AddCommand(textbots_bots_execute.Cmdtextbots_bots_execute())
	textbots_botsCmd.AddCommand(textbots_bots_search.Cmdtextbots_bots_search())
	textbots_botsCmd.Short = utils.GenerateCustomDescription(textbots_botsCmd.Short, textbots_bots_execute.Description, textbots_bots_search.Description, )
	textbots_botsCmd.Long = textbots_botsCmd.Short
}
