package analytics_botflows_divisions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_botflows_divisions", "SWAGGER_OVERRIDE_/api/v2/analytics/botflows/{botFlowId}/divisions")
	analytics_botflows_divisionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_botflows_divisions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_botflows_divisionsCmd)
}

func Cmdanalytics_botflows_divisions() *cobra.Command {
	return analytics_botflows_divisionsCmd
}
