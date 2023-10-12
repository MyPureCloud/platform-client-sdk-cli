package analytics_teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_teams", "SWAGGER_OVERRIDE_/api/v2/analytics/teams")
	analytics_teamsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_teams"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_teamsCmd)
}

func Cmdanalytics_teams() *cobra.Command {
	return analytics_teamsCmd
}
