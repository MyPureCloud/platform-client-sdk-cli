package conversations_messaging_oauth_apple

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_oauth_apple_callback"
)

func init() {
	conversations_messaging_oauth_appleCmd.AddCommand(conversations_messaging_oauth_apple_callback.Cmdconversations_messaging_oauth_apple_callback())
	conversations_messaging_oauth_appleCmd.Short = utils.GenerateCustomDescription(conversations_messaging_oauth_appleCmd.Short, conversations_messaging_oauth_apple_callback.Description, )
	conversations_messaging_oauth_appleCmd.Long = conversations_messaging_oauth_appleCmd.Short
}
