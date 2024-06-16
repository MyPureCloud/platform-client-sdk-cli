package learning

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assessments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_rules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_scheduleslots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_scorm"
)

func init() {
	learningCmd.AddCommand(learning_assessments.Cmdlearning_assessments())
	learningCmd.AddCommand(learning_assignments.Cmdlearning_assignments())
	learningCmd.AddCommand(learning_rules.Cmdlearning_rules())
	learningCmd.AddCommand(learning_modules.Cmdlearning_modules())
	learningCmd.AddCommand(learning_scheduleslots.Cmdlearning_scheduleslots())
	learningCmd.AddCommand(learning_scorm.Cmdlearning_scorm())
	learningCmd.Short = utils.GenerateCustomDescription(learningCmd.Short, learning_assessments.Description, learning_assignments.Description, learning_rules.Description, learning_modules.Description, learning_scheduleslots.Description, learning_scorm.Description, )
	learningCmd.Long = learningCmd.Short
}
