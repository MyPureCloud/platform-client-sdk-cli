package orgauthorization_trustors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustors_clonedusers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustors_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustors_users"
)

func init() {
	orgauthorization_trustorsCmd.AddCommand(orgauthorization_trustors_clonedusers.Cmdorgauthorization_trustors_clonedusers())
	orgauthorization_trustorsCmd.AddCommand(orgauthorization_trustors_groups.Cmdorgauthorization_trustors_groups())
	orgauthorization_trustorsCmd.AddCommand(orgauthorization_trustors_users.Cmdorgauthorization_trustors_users())
	orgauthorization_trustorsCmd.Short = utils.GenerateCustomDescription(orgauthorization_trustorsCmd.Short, orgauthorization_trustors_clonedusers.Description, orgauthorization_trustors_groups.Description, orgauthorization_trustors_users.Description, )
	orgauthorization_trustorsCmd.Long = orgauthorization_trustorsCmd.Short
}
