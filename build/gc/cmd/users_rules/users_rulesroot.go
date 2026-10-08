package users_rules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_rules_dependents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_rules_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_rules_settings"
)

func init() {
	users_rulesCmd.AddCommand(users_rules_dependents.Cmdusers_rules_dependents())
	users_rulesCmd.AddCommand(users_rules_query.Cmdusers_rules_query())
	users_rulesCmd.AddCommand(users_rules_settings.Cmdusers_rules_settings())
	users_rulesCmd.Short = utils.GenerateCustomDescription(users_rulesCmd.Short, users_rules_dependents.Description, users_rules_query.Description, users_rules_settings.Description, )
	users_rulesCmd.Long = users_rulesCmd.Short
}
