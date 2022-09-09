package analytics_dataretention

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_dataretention", "SWAGGER_OVERRIDE_/api/v2/analytics/dataretention")
	analytics_dataretentionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_dataretention"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_dataretentionCmd)
}

func Cmdanalytics_dataretention() *cobra.Command {
	return analytics_dataretentionCmd
}
