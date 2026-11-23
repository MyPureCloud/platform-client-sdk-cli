package analytics_agenticvirtualagents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agenticvirtualagents", "SWAGGER_OVERRIDE_/api/v2/analytics/agenticvirtualagents")
	analytics_agenticvirtualagentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agenticvirtualagents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agenticvirtualagentsCmd)
}

func Cmdanalytics_agenticvirtualagents() *cobra.Command {
	return analytics_agenticvirtualagentsCmd
}
