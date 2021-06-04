package orgauthorization_trustor

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustor_audits"
)

func init() {
	orgauthorization_trustorCmd.AddCommand(orgauthorization_trustor_audits.Cmdorgauthorization_trustor_audits())
	orgauthorization_trustorCmd.Short = utils.GenerateCustomDescription(orgauthorization_trustorCmd.Short, orgauthorization_trustor_audits.Description, )
	orgauthorization_trustorCmd.Long = orgauthorization_trustorCmd.Short
}
