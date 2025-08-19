package alerting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("alerting", "SWAGGER_OVERRIDE_/api/v2/alerting")
	alertingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("alerting"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(alertingCmd)
}

func Cmdalerting() *cobra.Command {
	return alertingCmd
}
