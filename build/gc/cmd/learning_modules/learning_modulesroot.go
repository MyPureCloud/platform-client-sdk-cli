package learning_modules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_rule"
)

func init() {
	learning_modulesCmd.AddCommand(learning_modules_versions.Cmdlearning_modules_versions())
	learning_modulesCmd.AddCommand(learning_modules_publish.Cmdlearning_modules_publish())
	learning_modulesCmd.AddCommand(learning_modules_rule.Cmdlearning_modules_rule())
	learning_modulesCmd.Short = utils.GenerateCustomDescription(learning_modulesCmd.Short, learning_modules_versions.Description, learning_modules_publish.Description, learning_modules_rule.Description, )
	learning_modulesCmd.Long = learning_modulesCmd.Short
}