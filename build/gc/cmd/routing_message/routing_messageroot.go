package routing_message

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_message_recipients"
)

func init() {
	routing_messageCmd.AddCommand(routing_message_recipients.Cmdrouting_message_recipients())
	routing_messageCmd.Short = utils.GenerateCustomDescription(routing_messageCmd.Short, routing_message_recipients.Description, )
	routing_messageCmd.Long = routing_messageCmd.Short
}
