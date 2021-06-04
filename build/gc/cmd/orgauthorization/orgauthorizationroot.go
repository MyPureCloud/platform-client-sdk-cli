package orgauthorization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustor"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustors"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_pairings"
)

func init() {
	orgauthorizationCmd.AddCommand(orgauthorization_trustor.Cmdorgauthorization_trustor())
	orgauthorizationCmd.AddCommand(orgauthorization_trustees.Cmdorgauthorization_trustees())
	orgauthorizationCmd.AddCommand(orgauthorization_trustors.Cmdorgauthorization_trustors())
	orgauthorizationCmd.AddCommand(orgauthorization_pairings.Cmdorgauthorization_pairings())
	orgauthorizationCmd.Short = utils.GenerateCustomDescription(orgauthorizationCmd.Short, orgauthorization_trustor.Description, orgauthorization_trustees.Description, orgauthorization_trustors.Description, orgauthorization_pairings.Description, )
	orgauthorizationCmd.Long = orgauthorizationCmd.Short
}
