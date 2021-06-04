package analytics_surveys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_surveys", "SWAGGER_OVERRIDE_/api/v2/analytics/surveys")
	analytics_surveysCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_surveys"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_surveysCmd)
}

func Cmdanalytics_surveys() *cobra.Command {
	return analytics_surveysCmd
}
