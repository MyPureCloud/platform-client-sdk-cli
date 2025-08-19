package conversations_messaging_supportedcontent

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_supportedcontent_default"
)

func init() {
	conversations_messaging_supportedcontentCmd.AddCommand(conversations_messaging_supportedcontent_default.Cmdconversations_messaging_supportedcontent_default())
	conversations_messaging_supportedcontentCmd.Short = utils.GenerateCustomDescription(conversations_messaging_supportedcontentCmd.Short, conversations_messaging_supportedcontent_default.Description, )
	conversations_messaging_supportedcontentCmd.Long = conversations_messaging_supportedcontentCmd.Short
}
