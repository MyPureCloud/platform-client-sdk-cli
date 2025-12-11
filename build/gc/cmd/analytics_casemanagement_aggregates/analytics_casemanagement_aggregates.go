package analytics_casemanagement_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_casemanagement_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/casemanagement/aggregates")
	analytics_casemanagement_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_casemanagement_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_casemanagement_aggregatesCmd)
}

func Cmdanalytics_casemanagement_aggregates() *cobra.Command {
	return analytics_casemanagement_aggregatesCmd
}
