package quality_evaluations_aggregates_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations_aggregates_query_me"
)

func init() {
	quality_evaluations_aggregates_queryCmd.AddCommand(quality_evaluations_aggregates_query_me.Cmdquality_evaluations_aggregates_query_me())
	quality_evaluations_aggregates_queryCmd.Short = utils.GenerateCustomDescription(quality_evaluations_aggregates_queryCmd.Short, quality_evaluations_aggregates_query_me.Description, )
	quality_evaluations_aggregates_queryCmd.Long = quality_evaluations_aggregates_queryCmd.Short
}
