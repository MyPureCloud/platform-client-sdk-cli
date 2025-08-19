package learning

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("learning", "SWAGGER_OVERRIDE_/api/v2/learning")
	learningCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learningCmd)
}

func Cmdlearning() *cobra.Command {
	return learningCmd
}
