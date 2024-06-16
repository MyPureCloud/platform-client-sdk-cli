package analytics_agentcopilots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agentcopilots", "SWAGGER_OVERRIDE_/api/v2/analytics/agentcopilots")
	analytics_agentcopilotsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agentcopilots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agentcopilotsCmd)
}

func Cmdanalytics_agentcopilots() *cobra.Command {
	return analytics_agentcopilotsCmd
}
