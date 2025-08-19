package textbots_botflows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("textbots_botflows", "SWAGGER_OVERRIDE_/api/v2/textbots/botflows")
	textbots_botflowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("textbots_botflows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(textbots_botflowsCmd)
}

func Cmdtextbots_botflows() *cobra.Command {
	return textbots_botflowsCmd
}
