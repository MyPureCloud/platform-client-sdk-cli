package telephony_sipmessages_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_sipmessages_conversations_headers"
)

func init() {
	telephony_sipmessages_conversationsCmd.AddCommand(telephony_sipmessages_conversations_headers.Cmdtelephony_sipmessages_conversations_headers())
	telephony_sipmessages_conversationsCmd.Short = utils.GenerateCustomDescription(telephony_sipmessages_conversationsCmd.Short, telephony_sipmessages_conversations_headers.Description, )
	telephony_sipmessages_conversationsCmd.Long = telephony_sipmessages_conversationsCmd.Short
}
