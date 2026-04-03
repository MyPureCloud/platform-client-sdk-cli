package analytics_copilots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_copilots", "SWAGGER_OVERRIDE_/api/v2/analytics/copilots")
	analytics_copilotsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_copilots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_copilotsCmd)
}

func Cmdanalytics_copilots() *cobra.Command {
	return analytics_copilotsCmd
}
