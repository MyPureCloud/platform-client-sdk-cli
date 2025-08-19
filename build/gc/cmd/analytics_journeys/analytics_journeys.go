package analytics_journeys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_journeys", "SWAGGER_OVERRIDE_/api/v2/analytics/journeys")
	analytics_journeysCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_journeys"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_journeysCmd)
}

func Cmdanalytics_journeys() *cobra.Command {
	return analytics_journeysCmd
}
