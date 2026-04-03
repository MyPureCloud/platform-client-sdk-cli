package conversations_calls_participants_snippet

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_snippet_record"
)

func init() {
	conversations_calls_participants_snippetCmd.AddCommand(conversations_calls_participants_snippet_record.Cmdconversations_calls_participants_snippet_record())
	conversations_calls_participants_snippetCmd.Short = utils.GenerateCustomDescription(conversations_calls_participants_snippetCmd.Short, conversations_calls_participants_snippet_record.Description, )
	conversations_calls_participants_snippetCmd.Long = conversations_calls_participants_snippetCmd.Short
}
