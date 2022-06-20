package conversations_screenshares

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_screenshares", "SWAGGER_OVERRIDE_/api/v2/conversations/screenshares")
	conversations_screensharesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_screenshares"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_screensharesCmd)
}

func Cmdconversations_screenshares() *cobra.Command {
	return conversations_screensharesCmd
}
