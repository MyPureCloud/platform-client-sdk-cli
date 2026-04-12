package analytics_agentutilizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agentutilizations", "SWAGGER_OVERRIDE_/api/v2/analytics/agentutilizations")
	analytics_agentutilizationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agentutilizations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agentutilizationsCmd)
}

func Cmdanalytics_agentutilizations() *cobra.Command {
	return analytics_agentutilizationsCmd
}
