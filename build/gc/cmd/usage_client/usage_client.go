package usage_client

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage_client", "SWAGGER_OVERRIDE_/api/v2/usage/client")
	usage_clientCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_client"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_clientCmd)
}

func Cmdusage_client() *cobra.Command {
	return usage_clientCmd
}
