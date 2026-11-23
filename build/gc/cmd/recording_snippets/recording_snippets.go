package recording_snippets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recording_snippets", "SWAGGER_OVERRIDE_/api/v2/recording/snippets")
	recording_snippetsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording_snippets"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recording_snippetsCmd)
}

func Cmdrecording_snippets() *cobra.Command {
	return recording_snippetsCmd
}
