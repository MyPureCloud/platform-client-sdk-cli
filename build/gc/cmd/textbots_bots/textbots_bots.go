package textbots_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("textbots_bots", "SWAGGER_OVERRIDE_/api/v2/textbots/bots")
	textbots_botsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("textbots_bots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(textbots_botsCmd)
}

func Cmdtextbots_bots() *cobra.Command {
	return textbots_botsCmd
}
