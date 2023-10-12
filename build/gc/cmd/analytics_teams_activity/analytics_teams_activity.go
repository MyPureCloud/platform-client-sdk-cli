package analytics_teams_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_teams_activity", "SWAGGER_OVERRIDE_/api/v2/analytics/teams/activity")
	analytics_teams_activityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_teams_activity"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_teams_activityCmd)
}

func Cmdanalytics_teams_activity() *cobra.Command {
	return analytics_teams_activityCmd
}
