package analytics_users_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_users_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/users/aggregates")
	analytics_users_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_users_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_users_aggregatesCmd)
}

func Cmdanalytics_users_aggregates() *cobra.Command {
	return analytics_users_aggregatesCmd
}
