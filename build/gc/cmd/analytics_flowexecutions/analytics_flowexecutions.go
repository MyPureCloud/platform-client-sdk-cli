package analytics_flowexecutions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_flowexecutions", "SWAGGER_OVERRIDE_/api/v2/analytics/flowexecutions")
	analytics_flowexecutionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_flowexecutions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_flowexecutionsCmd)
}

func Cmdanalytics_flowexecutions() *cobra.Command {
	return analytics_flowexecutionsCmd
}
