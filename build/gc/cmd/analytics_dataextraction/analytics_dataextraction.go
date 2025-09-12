package analytics_dataextraction

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_dataextraction", "SWAGGER_OVERRIDE_/api/v2/analytics/dataextraction")
	analytics_dataextractionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_dataextraction"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_dataextractionCmd)
}

func Cmdanalytics_dataextraction() *cobra.Command {
	return analytics_dataextractionCmd
}
