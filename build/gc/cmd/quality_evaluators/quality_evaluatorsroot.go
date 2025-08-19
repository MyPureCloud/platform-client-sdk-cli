package quality_evaluators

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_evaluators_activity"
)

func init() {
	quality_evaluatorsCmd.AddCommand(quality_evaluators_activity.Cmdquality_evaluators_activity())
	quality_evaluatorsCmd.Short = utils.GenerateCustomDescription(quality_evaluatorsCmd.Short, quality_evaluators_activity.Description, )
	quality_evaluatorsCmd.Long = quality_evaluatorsCmd.Short
}
