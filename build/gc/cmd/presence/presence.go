package presence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("presence", "SWAGGER_OVERRIDE_/api/v2/presence")
	presenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("presence"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(presenceCmd)
}

func Cmdpresence() *cobra.Command {
	return presenceCmd
}
