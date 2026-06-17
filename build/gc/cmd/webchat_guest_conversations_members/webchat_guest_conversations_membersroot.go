package webchat_guest_conversations_members

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest_conversations_members_typing"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest_conversations_members_messages"
)

func init() {
	webchat_guest_conversations_membersCmd.AddCommand(webchat_guest_conversations_members_typing.Cmdwebchat_guest_conversations_members_typing())
	webchat_guest_conversations_membersCmd.AddCommand(webchat_guest_conversations_members_messages.Cmdwebchat_guest_conversations_members_messages())
	webchat_guest_conversations_membersCmd.Short = utils.GenerateCustomDescription(webchat_guest_conversations_membersCmd.Short, webchat_guest_conversations_members_typing.Description, webchat_guest_conversations_members_messages.Description, )
	webchat_guest_conversations_membersCmd.Long = webchat_guest_conversations_membersCmd.Short
}
