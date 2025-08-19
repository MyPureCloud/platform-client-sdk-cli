package businessrules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_decisiontables"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/businessrules_schemas"
)

func init() {
	businessrulesCmd.AddCommand(businessrules_decisiontables.Cmdbusinessrules_decisiontables())
	businessrulesCmd.AddCommand(businessrules_schemas.Cmdbusinessrules_schemas())
	businessrulesCmd.Short = utils.GenerateCustomDescription(businessrulesCmd.Short, businessrules_decisiontables.Description, businessrules_schemas.Description, )
	businessrulesCmd.Long = businessrulesCmd.Short
}
