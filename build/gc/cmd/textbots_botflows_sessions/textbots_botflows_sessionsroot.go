package textbots_botflows_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_botflows_sessions_turns"
)

func init() {
	textbots_botflows_sessionsCmd.AddCommand(textbots_botflows_sessions_turns.Cmdtextbots_botflows_sessions_turns())
	textbots_botflows_sessionsCmd.Short = utils.GenerateCustomDescription(textbots_botflows_sessionsCmd.Short, textbots_botflows_sessions_turns.Description, )
	textbots_botflows_sessionsCmd.Long = textbots_botflows_sessionsCmd.Short
}
