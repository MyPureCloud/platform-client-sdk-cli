package analytics_summaries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_summaries", "SWAGGER_OVERRIDE_/api/v2/analytics/summaries")
	analytics_summariesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_summaries"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_summariesCmd)
}

func Cmdanalytics_summaries() *cobra.Command {
	return analytics_summariesCmd
}
