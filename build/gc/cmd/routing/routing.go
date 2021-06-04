package routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing", "SWAGGER_OVERRIDE_/api/v2/routing")
	routingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routingCmd)
}

func Cmdrouting() *cobra.Command {
	return routingCmd
}
