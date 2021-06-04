package users_development_activities_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_development_activities_aggregates", "SWAGGER_OVERRIDE_/api/v2/users/development/activities/aggregates")
	users_development_activities_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_development_activities_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_development_activities_aggregatesCmd)
}

func Cmdusers_development_activities_aggregates() *cobra.Command {
	return users_development_activities_aggregatesCmd
}
