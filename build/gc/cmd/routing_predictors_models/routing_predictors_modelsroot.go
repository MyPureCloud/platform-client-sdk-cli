package routing_predictors_models

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_predictors_models_features"
)

func init() {
	routing_predictors_modelsCmd.AddCommand(routing_predictors_models_features.Cmdrouting_predictors_models_features())
	routing_predictors_modelsCmd.Short = utils.GenerateCustomDescription(routing_predictors_modelsCmd.Short, routing_predictors_models_features.Description, )
	routing_predictors_modelsCmd.Long = routing_predictors_modelsCmd.Short
}
