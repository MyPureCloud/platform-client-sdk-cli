package analytics_surveys_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_surveys_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/surveys/aggregates")
	analytics_surveys_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_surveys_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_surveys_aggregatesCmd)
}

func Cmdanalytics_surveys_aggregates() *cobra.Command {
	return analytics_surveys_aggregatesCmd
}
