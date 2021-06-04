package conversations_messaging_facebook

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging_facebook", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/facebook")
	conversations_messaging_facebookCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_facebook"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_facebookCmd)
}

func Cmdconversations_messaging_facebook() *cobra.Command {
	return conversations_messaging_facebookCmd
}
