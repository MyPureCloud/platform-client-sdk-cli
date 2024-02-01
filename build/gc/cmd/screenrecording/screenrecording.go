package screenrecording

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("screenrecording", "SWAGGER_OVERRIDE_/api/v2/screenrecording")
	screenrecordingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("screenrecording"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(screenrecordingCmd)
}

func Cmdscreenrecording() *cobra.Command {
	return screenrecordingCmd
}
