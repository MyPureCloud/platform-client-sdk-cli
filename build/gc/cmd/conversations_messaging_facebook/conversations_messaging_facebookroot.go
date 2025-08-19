package conversations_messaging_facebook

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_facebook_app"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_facebook_permissions"
)

func init() {
	conversations_messaging_facebookCmd.AddCommand(conversations_messaging_facebook_app.Cmdconversations_messaging_facebook_app())
	conversations_messaging_facebookCmd.AddCommand(conversations_messaging_facebook_permissions.Cmdconversations_messaging_facebook_permissions())
	conversations_messaging_facebookCmd.Short = utils.GenerateCustomDescription(conversations_messaging_facebookCmd.Short, conversations_messaging_facebook_app.Description, conversations_messaging_facebook_permissions.Description, )
	conversations_messaging_facebookCmd.Long = conversations_messaging_facebookCmd.Short
}
