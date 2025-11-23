package learning_modules_rule

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_rule_migrate"
)

func init() {
	learning_modules_ruleCmd.AddCommand(learning_modules_rule_migrate.Cmdlearning_modules_rule_migrate())
	learning_modules_ruleCmd.Short = utils.GenerateCustomDescription(learning_modules_ruleCmd.Short, learning_modules_rule_migrate.Description, )
	learning_modules_ruleCmd.Long = learning_modules_ruleCmd.Short
}
