package organizations_limits_namespaces

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits_namespaces_defaults"
)

func init() {
	organizations_limits_namespacesCmd.AddCommand(organizations_limits_namespaces_defaults.Cmdorganizations_limits_namespaces_defaults())
	organizations_limits_namespacesCmd.Short = utils.GenerateCustomDescription(organizations_limits_namespacesCmd.Short, organizations_limits_namespaces_defaults.Description, )
	organizations_limits_namespacesCmd.Long = organizations_limits_namespacesCmd.Short
}
