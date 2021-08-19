package gamification_profiles_metrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_metrics_objectivedetails"
)

func init() {
	gamification_profiles_metricsCmd.AddCommand(gamification_profiles_metrics_objectivedetails.Cmdgamification_profiles_metrics_objectivedetails())
	gamification_profiles_metricsCmd.Short = utils.GenerateCustomDescription(gamification_profiles_metricsCmd.Short, gamification_profiles_metrics_objectivedetails.Description, )
	gamification_profiles_metricsCmd.Long = gamification_profiles_metricsCmd.Short
}
