package quality_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations_scoring"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluations_search"
)

func init() {
	quality_evaluationsCmd.AddCommand(quality_evaluations_aggregates.Cmdquality_evaluations_aggregates())
	quality_evaluationsCmd.AddCommand(quality_evaluations_query.Cmdquality_evaluations_query())
	quality_evaluationsCmd.AddCommand(quality_evaluations_scoring.Cmdquality_evaluations_scoring())
	quality_evaluationsCmd.AddCommand(quality_evaluations_search.Cmdquality_evaluations_search())
	quality_evaluationsCmd.Short = utils.GenerateCustomDescription(quality_evaluationsCmd.Short, quality_evaluations_aggregates.Description, quality_evaluations_query.Description, quality_evaluations_scoring.Description, quality_evaluations_search.Description, )
	quality_evaluationsCmd.Long = quality_evaluationsCmd.Short
}
