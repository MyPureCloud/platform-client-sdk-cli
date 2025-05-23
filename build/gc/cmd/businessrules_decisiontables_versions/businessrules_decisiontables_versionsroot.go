package businessrules_decisiontables_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions_copy"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions_execute"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions_publish"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions_sync"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions_rows"
)

func init() {
	businessrules_decisiontables_versionsCmd.AddCommand(businessrules_decisiontables_versions_copy.Cmdbusinessrules_decisiontables_versions_copy())
	businessrules_decisiontables_versionsCmd.AddCommand(businessrules_decisiontables_versions_execute.Cmdbusinessrules_decisiontables_versions_execute())
	businessrules_decisiontables_versionsCmd.AddCommand(businessrules_decisiontables_versions_publish.Cmdbusinessrules_decisiontables_versions_publish())
	businessrules_decisiontables_versionsCmd.AddCommand(businessrules_decisiontables_versions_sync.Cmdbusinessrules_decisiontables_versions_sync())
	businessrules_decisiontables_versionsCmd.AddCommand(businessrules_decisiontables_versions_rows.Cmdbusinessrules_decisiontables_versions_rows())
	businessrules_decisiontables_versionsCmd.Short = utils.GenerateCustomDescription(businessrules_decisiontables_versionsCmd.Short, businessrules_decisiontables_versions_copy.Description, businessrules_decisiontables_versions_execute.Description, businessrules_decisiontables_versions_publish.Description, businessrules_decisiontables_versions_sync.Description, businessrules_decisiontables_versions_rows.Description, )
	businessrules_decisiontables_versionsCmd.Long = businessrules_decisiontables_versionsCmd.Short
}
