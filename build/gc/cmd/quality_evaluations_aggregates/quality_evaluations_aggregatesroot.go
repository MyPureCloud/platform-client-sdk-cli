package quality_evaluations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations_aggregates_query"
)

func init() {
	quality_evaluations_aggregatesCmd.AddCommand(quality_evaluations_aggregates_query.Cmdquality_evaluations_aggregates_query())
	quality_evaluations_aggregatesCmd.Short = utils.GenerateCustomDescription(quality_evaluations_aggregatesCmd.Short, quality_evaluations_aggregates_query.Description, )
	quality_evaluations_aggregatesCmd.Long = quality_evaluations_aggregatesCmd.Short
}
