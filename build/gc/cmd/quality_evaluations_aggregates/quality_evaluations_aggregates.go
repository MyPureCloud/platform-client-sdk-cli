package quality_evaluations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_evaluations_aggregates", "SWAGGER_OVERRIDE_/api/v2/quality/evaluations/aggregates")
	quality_evaluations_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_evaluations_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_evaluations_aggregatesCmd)
}

func Cmdquality_evaluations_aggregates() *cobra.Command {
	return quality_evaluations_aggregatesCmd
}
