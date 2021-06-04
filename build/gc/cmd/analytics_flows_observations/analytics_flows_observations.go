package analytics_flows_observations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_flows_observations", "SWAGGER_OVERRIDE_/api/v2/analytics/flows/observations")
	analytics_flows_observationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_flows_observations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_flows_observationsCmd)
}

func Cmdanalytics_flows_observations() *cobra.Command {
	return analytics_flows_observationsCmd
}
