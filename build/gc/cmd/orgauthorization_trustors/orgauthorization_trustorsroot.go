package orgauthorization_trustors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustors_users"
)

func init() {
	orgauthorization_trustorsCmd.AddCommand(orgauthorization_trustors_users.Cmdorgauthorization_trustors_users())
	orgauthorization_trustorsCmd.Short = utils.GenerateCustomDescription(orgauthorization_trustorsCmd.Short, orgauthorization_trustors_users.Description, )
	orgauthorization_trustorsCmd.Long = orgauthorization_trustorsCmd.Short
}
