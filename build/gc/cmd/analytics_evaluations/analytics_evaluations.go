package analytics_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_evaluations", "SWAGGER_OVERRIDE_/api/v2/analytics/evaluations")
	analytics_evaluationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_evaluations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_evaluationsCmd)
}

func Cmdanalytics_evaluations() *cobra.Command {
	return analytics_evaluationsCmd
}
