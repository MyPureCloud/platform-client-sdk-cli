package businessrules_decisiontables

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_execute"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_exports"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_imports"
)

func init() {
	businessrules_decisiontablesCmd.AddCommand(businessrules_decisiontables_execute.Cmdbusinessrules_decisiontables_execute())
	businessrules_decisiontablesCmd.AddCommand(businessrules_decisiontables_versions.Cmdbusinessrules_decisiontables_versions())
	businessrules_decisiontablesCmd.AddCommand(businessrules_decisiontables_search.Cmdbusinessrules_decisiontables_search())
	businessrules_decisiontablesCmd.AddCommand(businessrules_decisiontables_exports.Cmdbusinessrules_decisiontables_exports())
	businessrules_decisiontablesCmd.AddCommand(businessrules_decisiontables_imports.Cmdbusinessrules_decisiontables_imports())
	businessrules_decisiontablesCmd.Short = utils.GenerateCustomDescription(businessrules_decisiontablesCmd.Short, businessrules_decisiontables_execute.Description, businessrules_decisiontables_versions.Description, businessrules_decisiontables_search.Description, businessrules_decisiontables_exports.Description, businessrules_decisiontables_imports.Description, )
	businessrules_decisiontablesCmd.Long = businessrules_decisiontablesCmd.Short
}
