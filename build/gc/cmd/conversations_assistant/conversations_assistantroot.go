package conversations_assistant

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_assistant_copilotcontext"
)

func init() {
	conversations_assistantCmd.AddCommand(conversations_assistant_copilotcontext.Cmdconversations_assistant_copilotcontext())
	conversations_assistantCmd.Short = utils.GenerateCustomDescription(conversations_assistantCmd.Short, conversations_assistant_copilotcontext.Description, )
	conversations_assistantCmd.Long = conversations_assistantCmd.Short
}
