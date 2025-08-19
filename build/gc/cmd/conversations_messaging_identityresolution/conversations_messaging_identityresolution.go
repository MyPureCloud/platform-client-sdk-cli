package conversations_messaging_identityresolution

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging_identityresolution", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/identityresolution")
	conversations_messaging_identityresolutionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_identityresolution"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_identityresolutionCmd)
}

func Cmdconversations_messaging_identityresolution() *cobra.Command {
	return conversations_messaging_identityresolutionCmd
}
