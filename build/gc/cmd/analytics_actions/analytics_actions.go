package analytics_actions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_actions", "SWAGGER_OVERRIDE_/api/v2/analytics/actions")
	analytics_actionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_actions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_actionsCmd)
}

func Cmdanalytics_actions() *cobra.Command {
	return analytics_actionsCmd
}
