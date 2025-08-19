package analytics_conversations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_conversations_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations/aggregates")
	analytics_conversations_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversations_aggregatesCmd)
}

func Cmdanalytics_conversations_aggregates() *cobra.Command {
	return analytics_conversations_aggregatesCmd
}
