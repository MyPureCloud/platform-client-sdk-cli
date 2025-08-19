package outbound_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_conversations_dnc"
)

func init() {
	outbound_conversationsCmd.AddCommand(outbound_conversations_dnc.Cmdoutbound_conversations_dnc())
	outbound_conversationsCmd.Short = utils.GenerateCustomDescription(outbound_conversationsCmd.Short, outbound_conversations_dnc.Description, )
	outbound_conversationsCmd.Long = outbound_conversationsCmd.Short
}
