package analytics_users_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_users_activity", "SWAGGER_OVERRIDE_/api/v2/analytics/users/activity")
	analytics_users_activityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_users_activity"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_users_activityCmd)
}

func Cmdanalytics_users_activity() *cobra.Command {
	return analytics_users_activityCmd
}
