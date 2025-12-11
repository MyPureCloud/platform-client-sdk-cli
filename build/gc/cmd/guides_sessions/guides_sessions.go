package guides_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("guides_sessions", "SWAGGER_OVERRIDE_/api/v2/guides/{guideId}/sessions")
	guides_sessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("guides_sessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(guides_sessionsCmd)
}

func Cmdguides_sessions() *cobra.Command {
	return guides_sessionsCmd
}
