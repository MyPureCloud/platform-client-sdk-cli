package workforcemanagement_teams_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_teams_adherence", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/teams/{teamId}/adherence")
	workforcemanagement_teams_adherenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_teams_adherence"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_teams_adherenceCmd)
}

func Cmdworkforcemanagement_teams_adherence() *cobra.Command {
	return workforcemanagement_teams_adherenceCmd
}
