package analytics_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_conversations", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations")
	analytics_conversationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversationsCmd)
}

func Cmdanalytics_conversations() *cobra.Command {
	return analytics_conversationsCmd
}
