package learning

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_rules"
)

func init() {
	learningCmd.AddCommand(learning_assignments.Cmdlearning_assignments())
	learningCmd.AddCommand(learning_modules.Cmdlearning_modules())
	learningCmd.AddCommand(learning_rules.Cmdlearning_rules())
	learningCmd.Short = utils.GenerateCustomDescription(learningCmd.Short, learning_assignments.Description, learning_modules.Description, learning_rules.Description, )
	learningCmd.Long = learningCmd.Short
}
