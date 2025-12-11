package guides_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_sessions_turns"
)

func init() {
	guides_sessionsCmd.AddCommand(guides_sessions_turns.Cmdguides_sessions_turns())
	guides_sessionsCmd.Short = utils.GenerateCustomDescription(guides_sessionsCmd.Short, guides_sessions_turns.Description, )
	guides_sessionsCmd.Long = guides_sessionsCmd.Short
}
