package alerting_interactionstats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("alerting_interactionstats", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats")
	alerting_interactionstatsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("alerting_interactionstats"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(alerting_interactionstatsCmd)
}

func Cmdalerting_interactionstats() *cobra.Command {
	return alerting_interactionstatsCmd
}
