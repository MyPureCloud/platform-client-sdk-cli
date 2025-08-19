package textbots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("textbots", "SWAGGER_OVERRIDE_/api/v2/textbots")
	textbotsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("textbots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(textbotsCmd)
}

func Cmdtextbots() *cobra.Command {
	return textbotsCmd
}
