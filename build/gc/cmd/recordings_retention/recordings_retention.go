package recordings_retention

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recordings_retention", "SWAGGER_OVERRIDE_/api/v2/recordings/retention")
	recordings_retentionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recordings_retention"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recordings_retentionCmd)
}

func Cmdrecordings_retention() *cobra.Command {
	return recordings_retentionCmd
}
