package conversations_calls

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_maximumconferenceparties"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_history"
)

func init() {
	conversations_callsCmd.AddCommand(conversations_calls_participants.Cmdconversations_calls_participants())
	conversations_callsCmd.AddCommand(conversations_calls_maximumconferenceparties.Cmdconversations_calls_maximumconferenceparties())
	conversations_callsCmd.AddCommand(conversations_calls_history.Cmdconversations_calls_history())
	conversations_callsCmd.Short = utils.GenerateCustomDescription(conversations_callsCmd.Short, conversations_calls_participants.Description, conversations_calls_maximumconferenceparties.Description, conversations_calls_history.Description, )
	conversations_callsCmd.Long = conversations_callsCmd.Short
}
