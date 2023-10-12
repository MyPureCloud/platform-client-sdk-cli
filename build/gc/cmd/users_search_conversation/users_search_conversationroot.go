package users_search_conversation

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_search_conversation_target"
)

func init() {
	users_search_conversationCmd.AddCommand(users_search_conversation_target.Cmdusers_search_conversation_target())
	users_search_conversationCmd.Short = utils.GenerateCustomDescription(users_search_conversationCmd.Short, users_search_conversation_target.Description, )
	users_search_conversationCmd.Long = users_search_conversationCmd.Short
}
