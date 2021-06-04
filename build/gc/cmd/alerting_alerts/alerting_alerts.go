package alerting_alerts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("alerting_alerts", "SWAGGER_OVERRIDE_/api/v2/alerting/alerts")
	alerting_alertsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("alerting_alerts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(alerting_alertsCmd)
}

func Cmdalerting_alerts() *cobra.Command {
	return alerting_alertsCmd
}
