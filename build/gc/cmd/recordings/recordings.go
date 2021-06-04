package recordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recordings", "SWAGGER_OVERRIDE_/api/v2/recordings")
	recordingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recordings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recordingsCmd)
}

func Cmdrecordings() *cobra.Command {
	return recordingsCmd
}
