package socialmedia_escalations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("socialmedia_escalations", "SWAGGER_OVERRIDE_/api/v2/socialmedia/escalations")
	socialmedia_escalationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("socialmedia_escalations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(socialmedia_escalationsCmd)
}

func Cmdsocialmedia_escalations() *cobra.Command {
	return socialmedia_escalationsCmd
}
