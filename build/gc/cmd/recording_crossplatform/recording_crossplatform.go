package recording_crossplatform

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recording_crossplatform", "SWAGGER_OVERRIDE_/api/v2/recording/crossplatform")
	recording_crossplatformCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording_crossplatform"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recording_crossplatformCmd)
}

func Cmdrecording_crossplatform() *cobra.Command {
	return recording_crossplatformCmd
}
