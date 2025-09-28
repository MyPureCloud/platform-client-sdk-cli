package conversations_messaging_oauth

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_oauth_apple"
)

func init() {
	conversations_messaging_oauthCmd.AddCommand(conversations_messaging_oauth_apple.Cmdconversations_messaging_oauth_apple())
	conversations_messaging_oauthCmd.Short = utils.GenerateCustomDescription(conversations_messaging_oauthCmd.Short, conversations_messaging_oauth_apple.Description, )
	conversations_messaging_oauthCmd.Long = conversations_messaging_oauthCmd.Short
}
