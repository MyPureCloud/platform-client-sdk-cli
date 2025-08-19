package analytics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics", "SWAGGER_OVERRIDE_/api/v2/analytics")
	analyticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analyticsCmd)
}

func Cmdanalytics() *cobra.Command {
	return analyticsCmd
}
