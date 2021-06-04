package organizations_limits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits_changerequests"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits_namespaces"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits_docs"
)

func init() {
	organizations_limitsCmd.AddCommand(organizations_limits_changerequests.Cmdorganizations_limits_changerequests())
	organizations_limitsCmd.AddCommand(organizations_limits_namespaces.Cmdorganizations_limits_namespaces())
	organizations_limitsCmd.AddCommand(organizations_limits_docs.Cmdorganizations_limits_docs())
	organizations_limitsCmd.Short = utils.GenerateCustomDescription(organizations_limitsCmd.Short, organizations_limits_changerequests.Description, organizations_limits_namespaces.Description, organizations_limits_docs.Description, )
	organizations_limitsCmd.Long = organizations_limitsCmd.Short
}
