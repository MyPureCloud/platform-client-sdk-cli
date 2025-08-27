package assistants_queues_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("assistants_queues_users", "SWAGGER_OVERRIDE_/api/v2/assistants/{assistantId}/queues/{queueId}/users")
	assistants_queues_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("assistants_queues_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(assistants_queues_usersCmd)
}

func Cmdassistants_queues_users() *cobra.Command {
	return assistants_queues_usersCmd
}
