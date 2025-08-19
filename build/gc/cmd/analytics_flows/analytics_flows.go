package analytics_flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_flows", "SWAGGER_OVERRIDE_/api/v2/analytics/flows")
	analytics_flowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_flows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_flowsCmd)
}

func Cmdanalytics_flows() *cobra.Command {
	return analytics_flowsCmd
}
