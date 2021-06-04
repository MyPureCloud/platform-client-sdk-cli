package usage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage", "SWAGGER_OVERRIDE_/api/v2/usage")
	usageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usageCmd)
}

func Cmdusage() *cobra.Command {
	return usageCmd
}
