package analytics_routing_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_routing_activity", "SWAGGER_OVERRIDE_/api/v2/analytics/routing/activity")
	analytics_routing_activityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_routing_activity"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_routing_activityCmd)
}

func Cmdanalytics_routing_activity() *cobra.Command {
	return analytics_routing_activityCmd
}
