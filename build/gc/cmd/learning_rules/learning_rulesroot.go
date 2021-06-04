package learning_rules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_rules_query"
)

func init() {
	learning_rulesCmd.AddCommand(learning_rules_query.Cmdlearning_rules_query())
	learning_rulesCmd.Short = utils.GenerateCustomDescription(learning_rulesCmd.Short, learning_rules_query.Description, )
	learning_rulesCmd.Long = learning_rulesCmd.Short
}
