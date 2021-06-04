package webchat_guest

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("webchat_guest", "SWAGGER_OVERRIDE_/api/v2/webchat/guest")
	webchat_guestCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat_guest"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchat_guestCmd)
}

func Cmdwebchat_guest() *cobra.Command {
	return webchat_guestCmd
}
