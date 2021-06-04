package quality_evaluators

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_evaluators", "SWAGGER_OVERRIDE_/api/v2/quality/evaluators")
	quality_evaluatorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_evaluators"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_evaluatorsCmd)
}

func Cmdquality_evaluators() *cobra.Command {
	return quality_evaluatorsCmd
}
