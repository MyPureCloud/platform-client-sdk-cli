package users_search_teams

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_search_teams", "SWAGGER_OVERRIDE_/api/v2/users/search/teams")
	users_search_teamsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_search_teams"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_search_teamsCmd)
}

func Cmdusers_search_teams() *cobra.Command {
	return users_search_teamsCmd
}
