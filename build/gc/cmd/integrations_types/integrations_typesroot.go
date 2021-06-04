package integrations_types

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_types_configschemas"
)

func init() {
	integrations_typesCmd.AddCommand(integrations_types_configschemas.Cmdintegrations_types_configschemas())
	integrations_typesCmd.Short = utils.GenerateCustomDescription(integrations_typesCmd.Short, integrations_types_configschemas.Description, )
	integrations_typesCmd.Long = integrations_typesCmd.Short
}
