package webchat

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("webchat", "SWAGGER_OVERRIDE_/api/v2/webchat")
	webchatCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchatCmd)
}

func Cmdwebchat() *cobra.Command {
	return webchatCmd
}
