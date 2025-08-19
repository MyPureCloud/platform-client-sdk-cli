package quality

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality", "SWAGGER_OVERRIDE_/api/v2/quality")
	qualityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(qualityCmd)
}

func Cmdquality() *cobra.Command {
	return qualityCmd
}
