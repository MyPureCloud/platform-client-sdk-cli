package quality_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_evaluations", "SWAGGER_OVERRIDE_/api/v2/quality/evaluations")
	quality_evaluationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_evaluations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_evaluationsCmd)
}

func Cmdquality_evaluations() *cobra.Command {
	return quality_evaluationsCmd
}
