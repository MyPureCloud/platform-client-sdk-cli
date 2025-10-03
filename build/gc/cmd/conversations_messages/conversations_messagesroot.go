package conversations_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_cachedmedia"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_communications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_agentless"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_inbound"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_recordingstate"
)

func init() {
	conversations_messagesCmd.AddCommand(conversations_messages_cachedmedia.Cmdconversations_messages_cachedmedia())
	conversations_messagesCmd.AddCommand(conversations_messages_participants.Cmdconversations_messages_participants())
	conversations_messagesCmd.AddCommand(conversations_messages_details.Cmdconversations_messages_details())
	conversations_messagesCmd.AddCommand(conversations_messages_communications.Cmdconversations_messages_communications())
	conversations_messagesCmd.AddCommand(conversations_messages_messages.Cmdconversations_messages_messages())
	conversations_messagesCmd.AddCommand(conversations_messages_agentless.Cmdconversations_messages_agentless())
	conversations_messagesCmd.AddCommand(conversations_messages_inbound.Cmdconversations_messages_inbound())
	conversations_messagesCmd.AddCommand(conversations_messages_recordingstate.Cmdconversations_messages_recordingstate())
	conversations_messagesCmd.Short = utils.GenerateCustomDescription(conversations_messagesCmd.Short, conversations_messages_cachedmedia.Description, conversations_messages_participants.Description, conversations_messages_details.Description, conversations_messages_communications.Description, conversations_messages_messages.Description, conversations_messages_agentless.Description, conversations_messages_inbound.Description, conversations_messages_recordingstate.Description, )
	conversations_messagesCmd.Long = conversations_messagesCmd.Short
}
