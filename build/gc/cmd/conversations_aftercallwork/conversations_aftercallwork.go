package conversations_aftercallwork

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_aftercallwork", "SWAGGER_OVERRIDE_/api/v2/conversations/aftercallwork")
	conversations_aftercallworkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_aftercallwork"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_aftercallworkCmd)
}

func Cmdconversations_aftercallwork() *cobra.Command {
	return conversations_aftercallworkCmd
}
