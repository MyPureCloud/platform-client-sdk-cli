package knowledge_guest

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_guest", "SWAGGER_OVERRIDE_/api/v2/knowledge/guest")
	knowledge_guestCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_guest"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_guestCmd)
}

func Cmdknowledge_guest() *cobra.Command {
	return knowledge_guestCmd
}
