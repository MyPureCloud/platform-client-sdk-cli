package recording

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recording", "SWAGGER_OVERRIDE_/api/v2/recording")
	recordingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recordingCmd)
}

func Cmdrecording() *cobra.Command {
	return recordingCmd
}
