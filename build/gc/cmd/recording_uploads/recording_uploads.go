package recording_uploads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("recording_uploads", "SWAGGER_OVERRIDE_/api/v2/recording/uploads")
	recording_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("recording_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(recording_uploadsCmd)
}

func Cmdrecording_uploads() *cobra.Command {
	return recording_uploadsCmd
}
