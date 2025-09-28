package conversations_messaging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_integrations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_threadingtimeline"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_identityresolution"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_facebook"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_supportedcontent"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_oauth"
)

func init() {
	conversations_messagingCmd.AddCommand(conversations_messaging_integrations.Cmdconversations_messaging_integrations())
	conversations_messagingCmd.AddCommand(conversations_messaging_threadingtimeline.Cmdconversations_messaging_threadingtimeline())
	conversations_messagingCmd.AddCommand(conversations_messaging_identityresolution.Cmdconversations_messaging_identityresolution())
	conversations_messagingCmd.AddCommand(conversations_messaging_facebook.Cmdconversations_messaging_facebook())
	conversations_messagingCmd.AddCommand(conversations_messaging_settings.Cmdconversations_messaging_settings())
	conversations_messagingCmd.AddCommand(conversations_messaging_supportedcontent.Cmdconversations_messaging_supportedcontent())
	conversations_messagingCmd.AddCommand(conversations_messaging_oauth.Cmdconversations_messaging_oauth())
	conversations_messagingCmd.Short = utils.GenerateCustomDescription(conversations_messagingCmd.Short, conversations_messaging_integrations.Description, conversations_messaging_threadingtimeline.Description, conversations_messaging_identityresolution.Description, conversations_messaging_facebook.Description, conversations_messaging_settings.Description, conversations_messaging_supportedcontent.Description, conversations_messaging_oauth.Description, )
	conversations_messagingCmd.Long = conversations_messagingCmd.Short
}
