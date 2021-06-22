package webchat_guest_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest_conversations_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest_conversations_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest_conversations_mediarequests"
)

func init() {
	webchat_guest_conversationsCmd.AddCommand(webchat_guest_conversations_messages.Cmdwebchat_guest_conversations_messages())
	webchat_guest_conversationsCmd.AddCommand(webchat_guest_conversations_members.Cmdwebchat_guest_conversations_members())
	webchat_guest_conversationsCmd.AddCommand(webchat_guest_conversations_mediarequests.Cmdwebchat_guest_conversations_mediarequests())
	webchat_guest_conversationsCmd.Short = utils.GenerateCustomDescription(webchat_guest_conversationsCmd.Short, webchat_guest_conversations_messages.Description, webchat_guest_conversations_members.Description, webchat_guest_conversations_mediarequests.Description, )
	webchat_guest_conversationsCmd.Long = webchat_guest_conversationsCmd.Short
}
