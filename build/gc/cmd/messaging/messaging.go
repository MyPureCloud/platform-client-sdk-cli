package messaging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("messaging", "SWAGGER_OVERRIDE_/api/v2/messaging")
	messagingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("messaging"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(messagingCmd)
}

func Cmdmessaging() *cobra.Command {
	return messagingCmd
}
