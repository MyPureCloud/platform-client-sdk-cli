package organizations_limits_docs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits_docs_freetrial"
)

func init() {
	organizations_limits_docsCmd.AddCommand(organizations_limits_docs_freetrial.Cmdorganizations_limits_docs_freetrial())
	organizations_limits_docsCmd.Short = utils.GenerateCustomDescription(organizations_limits_docsCmd.Short, organizations_limits_docs_freetrial.Description, )
	organizations_limits_docsCmd.Long = organizations_limits_docsCmd.Short
}
