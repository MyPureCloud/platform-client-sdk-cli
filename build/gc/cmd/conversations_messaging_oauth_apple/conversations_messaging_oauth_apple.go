package conversations_messaging_oauth_apple

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging_oauth_apple", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/oauth/apple")
	conversations_messaging_oauth_appleCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_oauth_apple"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_oauth_appleCmd)
}

func Cmdconversations_messaging_oauth_apple() *cobra.Command {
	return conversations_messaging_oauth_appleCmd
}
