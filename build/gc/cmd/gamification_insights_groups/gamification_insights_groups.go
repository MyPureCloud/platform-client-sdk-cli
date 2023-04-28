package gamification_insights_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_insights_groups", "SWAGGER_OVERRIDE_/api/v2/gamification/insights/groups")
	gamification_insights_groupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_insights_groups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_insights_groupsCmd)
}

func Cmdgamification_insights_groups() *cobra.Command {
	return gamification_insights_groupsCmd
}
