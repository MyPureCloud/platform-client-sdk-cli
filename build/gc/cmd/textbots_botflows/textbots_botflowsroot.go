package textbots_botflows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_botflows_sessions"
)

func init() {
	textbots_botflowsCmd.AddCommand(textbots_botflows_sessions.Cmdtextbots_botflows_sessions())
	textbots_botflowsCmd.Short = utils.GenerateCustomDescription(textbots_botflowsCmd.Short, textbots_botflows_sessions.Description, )
	textbots_botflowsCmd.Long = textbots_botflowsCmd.Short
}
