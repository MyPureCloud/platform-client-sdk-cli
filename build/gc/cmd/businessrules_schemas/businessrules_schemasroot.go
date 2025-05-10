package businessrules_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_schemas_coretypes"
)

func init() {
	businessrules_schemasCmd.AddCommand(businessrules_schemas_coretypes.Cmdbusinessrules_schemas_coretypes())
	businessrules_schemasCmd.Short = utils.GenerateCustomDescription(businessrules_schemasCmd.Short, businessrules_schemas_coretypes.Description, )
	businessrules_schemasCmd.Long = businessrules_schemasCmd.Short
}
