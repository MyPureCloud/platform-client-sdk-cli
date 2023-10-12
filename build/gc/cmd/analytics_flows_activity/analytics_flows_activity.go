package analytics_flows_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_flows_activity", "SWAGGER_OVERRIDE_/api/v2/analytics/flows/activity")
	analytics_flows_activityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_flows_activity"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_flows_activityCmd)
}

func Cmdanalytics_flows_activity() *cobra.Command {
	return analytics_flows_activityCmd
}
