package quality_evaluations_aggregates_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_evaluations_aggregates_query", "SWAGGER_OVERRIDE_/api/v2/quality/evaluations/aggregates/query")
	quality_evaluations_aggregates_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_evaluations_aggregates_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_evaluations_aggregates_queryCmd)
}

func Cmdquality_evaluations_aggregates_query() *cobra.Command {
	return quality_evaluations_aggregates_queryCmd
}
