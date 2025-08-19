package workforcemanagement_teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_teams", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/teams")
	workforcemanagement_teamsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_teams"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_teamsCmd)
}

func Cmdworkforcemanagement_teams() *cobra.Command {
	return workforcemanagement_teamsCmd
}
