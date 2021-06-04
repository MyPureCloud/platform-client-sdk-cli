package conversations_messages_inbound

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_inbound_open"
)

func init() {
	conversations_messages_inboundCmd.AddCommand(conversations_messages_inbound_open.Cmdconversations_messages_inbound_open())
	conversations_messages_inboundCmd.Short = utils.GenerateCustomDescription(conversations_messages_inboundCmd.Short, conversations_messages_inbound_open.Description, )
	conversations_messages_inboundCmd.Long = conversations_messages_inboundCmd.Short
}
