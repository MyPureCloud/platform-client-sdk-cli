package conversations_callbacks_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_callbacks_bulk", "SWAGGER_OVERRIDE_/api/v2/conversations/callbacks/bulk")
	conversations_callbacks_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_callbacks_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_callbacks_bulkCmd)
}

func Cmdconversations_callbacks_bulk() *cobra.Command {
	return conversations_callbacks_bulkCmd
}
