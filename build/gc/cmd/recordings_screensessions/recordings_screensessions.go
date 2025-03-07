package recordings_screensessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recordings_screensessions", "SWAGGER_OVERRIDE_/api/v2/recordings/screensessions")
	recordings_screensessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recordings_screensessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recordings_screensessionsCmd)
}

func Cmdrecordings_screensessions() *cobra.Command {
	return recordings_screensessionsCmd
}
