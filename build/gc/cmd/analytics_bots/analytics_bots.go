package analytics_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_bots", "SWAGGER_OVERRIDE_/api/v2/analytics/bots")
	analytics_botsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_bots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_botsCmd)
}

func Cmdanalytics_bots() *cobra.Command {
	return analytics_botsCmd
}
