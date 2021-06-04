package analytics_users_observations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_users_observations", "SWAGGER_OVERRIDE_/api/v2/analytics/users/observations")
	analytics_users_observationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_users_observations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_users_observationsCmd)
}

func Cmdanalytics_users_observations() *cobra.Command {
	return analytics_users_observationsCmd
}
