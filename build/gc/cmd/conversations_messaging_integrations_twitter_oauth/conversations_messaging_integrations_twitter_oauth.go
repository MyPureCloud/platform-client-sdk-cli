package conversations_messaging_integrations_twitter_oauth

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging_integrations_twitter_oauth", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/integrations/twitter/oauth")
	conversations_messaging_integrations_twitter_oauthCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_integrations_twitter_oauth"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_integrations_twitter_oauthCmd)
}

func Cmdconversations_messaging_integrations_twitter_oauth() *cobra.Command {
	return conversations_messaging_integrations_twitter_oauthCmd
}
