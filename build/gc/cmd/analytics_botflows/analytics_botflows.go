package analytics_botflows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_botflows", "SWAGGER_OVERRIDE_/api/v2/analytics/botflows")
	analytics_botflowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_botflows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_botflowsCmd)
}

func Cmdanalytics_botflows() *cobra.Command {
	return analytics_botflowsCmd
}
