package orgauthorization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_pairings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustors"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustor"
)

func init() {
	orgauthorizationCmd.AddCommand(orgauthorization_pairings.Cmdorgauthorization_pairings())
	orgauthorizationCmd.AddCommand(orgauthorization_trustees.Cmdorgauthorization_trustees())
	orgauthorizationCmd.AddCommand(orgauthorization_trustors.Cmdorgauthorization_trustors())
	orgauthorizationCmd.AddCommand(orgauthorization_trustor.Cmdorgauthorization_trustor())
	orgauthorizationCmd.Short = utils.GenerateCustomDescription(orgauthorizationCmd.Short, orgauthorization_pairings.Description, orgauthorization_trustees.Description, orgauthorization_trustors.Description, orgauthorization_trustor.Description, )
	orgauthorizationCmd.Long = orgauthorizationCmd.Short
}
