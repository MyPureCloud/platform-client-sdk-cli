package workforcemanagement_teams_shrinkage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_teams_shrinkage", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/teams/{teamId}/shrinkage")
	workforcemanagement_teams_shrinkageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_teams_shrinkage"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_teams_shrinkageCmd)
}

func Cmdworkforcemanagement_teams_shrinkage() *cobra.Command {
	return workforcemanagement_teams_shrinkageCmd
}
