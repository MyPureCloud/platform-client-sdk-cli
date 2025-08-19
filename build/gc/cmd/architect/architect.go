package architect

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("architect", "SWAGGER_OVERRIDE_/api/v2/architect")
	architectCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architectCmd)
}

func Cmdarchitect() *cobra.Command {
	return architectCmd
}
