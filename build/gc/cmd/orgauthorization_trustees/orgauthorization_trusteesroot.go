package orgauthorization_trustees

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees_audits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees_users"
)

func init() {
	orgauthorization_trusteesCmd.AddCommand(orgauthorization_trustees_audits.Cmdorgauthorization_trustees_audits())
	orgauthorization_trusteesCmd.AddCommand(orgauthorization_trustees_users.Cmdorgauthorization_trustees_users())
	orgauthorization_trusteesCmd.Short = utils.GenerateCustomDescription(orgauthorization_trusteesCmd.Short, orgauthorization_trustees_audits.Description, orgauthorization_trustees_users.Description, )
	orgauthorization_trusteesCmd.Long = orgauthorization_trusteesCmd.Short
}
