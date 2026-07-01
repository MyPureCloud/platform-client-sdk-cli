package screenmonitors_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("screenmonitors_sessions", "SWAGGER_OVERRIDE_/api/v2/screenmonitors/sessions")
	screenmonitors_sessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("screenmonitors_sessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(screenmonitors_sessionsCmd)
}

func Cmdscreenmonitors_sessions() *cobra.Command {
	return screenmonitors_sessionsCmd
}
