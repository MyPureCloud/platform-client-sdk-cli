package assistants_queues_users_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("assistants_queues_users_bulk", "SWAGGER_OVERRIDE_/api/v2/assistants/{assistantId}/queues/{queueId}/users/bulk")
	assistants_queues_users_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("assistants_queues_users_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(assistants_queues_users_bulkCmd)
}

func Cmdassistants_queues_users_bulk() *cobra.Command {
	return assistants_queues_users_bulkCmd
}
