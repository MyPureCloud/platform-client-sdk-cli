package businessrules_decisiontables_versions_rows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables_versions_rows_search"
)

func init() {
	businessrules_decisiontables_versions_rowsCmd.AddCommand(businessrules_decisiontables_versions_rows_search.Cmdbusinessrules_decisiontables_versions_rows_search())
	businessrules_decisiontables_versions_rowsCmd.Short = utils.GenerateCustomDescription(businessrules_decisiontables_versions_rowsCmd.Short, businessrules_decisiontables_versions_rows_search.Description, )
	businessrules_decisiontables_versions_rowsCmd.Long = businessrules_decisiontables_versions_rowsCmd.Short
}
