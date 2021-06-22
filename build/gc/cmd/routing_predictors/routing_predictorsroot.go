package routing_predictors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_predictors_keyperformanceindicators"
)

func init() {
	routing_predictorsCmd.AddCommand(routing_predictors_keyperformanceindicators.Cmdrouting_predictors_keyperformanceindicators())
	routing_predictorsCmd.Short = utils.GenerateCustomDescription(routing_predictorsCmd.Short, routing_predictors_keyperformanceindicators.Description, )
	routing_predictorsCmd.Long = routing_predictorsCmd.Short
}
