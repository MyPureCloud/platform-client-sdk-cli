package events_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("events_users", "SWAGGER_OVERRIDE_/api/v2/events/users")
	events_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("events_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(events_usersCmd)
}

func Cmdevents_users() *cobra.Command {
	return events_usersCmd
}
