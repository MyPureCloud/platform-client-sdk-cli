package analytics_conversations_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_conversations_activity", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations/activity")
	analytics_conversations_activityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations_activity"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversations_activityCmd)
}

func Cmdanalytics_conversations_activity() *cobra.Command {
	return analytics_conversations_activityCmd
}
