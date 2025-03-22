package socialmedia_analytics_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_analytics_messages", "SWAGGER_OVERRIDE_/api/v2/socialmedia/analytics/messages")
	socialmedia_analytics_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_analytics_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_analytics_messagesCmd)
}

func Cmdsocialmedia_analytics_messages() *cobra.Command {
	return socialmedia_analytics_messagesCmd
}
