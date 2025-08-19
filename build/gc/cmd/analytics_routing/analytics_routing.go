package analytics_routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_routing", "SWAGGER_OVERRIDE_/api/v2/analytics/routing")
	analytics_routingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_routing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_routingCmd)
}

func Cmdanalytics_routing() *cobra.Command {
	return analytics_routingCmd
}
