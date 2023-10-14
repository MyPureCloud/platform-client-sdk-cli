package telephony_sipmessages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_sipmessages_conversations"
)

func init() {
	telephony_sipmessagesCmd.AddCommand(telephony_sipmessages_conversations.Cmdtelephony_sipmessages_conversations())
	telephony_sipmessagesCmd.Short = utils.GenerateCustomDescription(telephony_sipmessagesCmd.Short, telephony_sipmessages_conversations.Description, )
	telephony_sipmessagesCmd.Long = telephony_sipmessagesCmd.Short
}
