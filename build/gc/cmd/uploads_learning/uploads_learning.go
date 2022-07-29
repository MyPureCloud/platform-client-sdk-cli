package uploads_learning

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("uploads_learning", "SWAGGER_OVERRIDE_/api/v2/uploads/learning")
	uploads_learningCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("uploads_learning"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(uploads_learningCmd)
}

func Cmduploads_learning() *cobra.Command {
	return uploads_learningCmd
}
