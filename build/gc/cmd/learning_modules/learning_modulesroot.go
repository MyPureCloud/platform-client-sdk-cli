package learning_modules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_coverart"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_assignments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_rule"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_users"
)

func init() {
	learning_modulesCmd.AddCommand(learning_modules_coverart.Cmdlearning_modules_coverart())
	learning_modulesCmd.AddCommand(learning_modules_jobs.Cmdlearning_modules_jobs())
	learning_modulesCmd.AddCommand(learning_modules_versions.Cmdlearning_modules_versions())
	learning_modulesCmd.AddCommand(learning_modules_publish.Cmdlearning_modules_publish())
	learning_modulesCmd.AddCommand(learning_modules_assignments.Cmdlearning_modules_assignments())
	learning_modulesCmd.AddCommand(learning_modules_rule.Cmdlearning_modules_rule())
	learning_modulesCmd.AddCommand(learning_modules_users.Cmdlearning_modules_users())
	learning_modulesCmd.Short = utils.GenerateCustomDescription(learning_modulesCmd.Short, learning_modules_coverart.Description, learning_modules_jobs.Description, learning_modules_versions.Description, learning_modules_publish.Description, learning_modules_assignments.Description, learning_modules_rule.Description, learning_modules_users.Description, )
	learning_modulesCmd.Long = learning_modulesCmd.Short
}
