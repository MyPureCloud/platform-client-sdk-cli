package organizations_limits_namespaces_limits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations_limits_namespaces_limits_counts"
)

func init() {
	organizations_limits_namespaces_limitsCmd.AddCommand(organizations_limits_namespaces_limits_counts.Cmdorganizations_limits_namespaces_limits_counts())
	organizations_limits_namespaces_limitsCmd.Short = utils.GenerateCustomDescription(organizations_limits_namespaces_limitsCmd.Short, organizations_limits_namespaces_limits_counts.Description, )
	organizations_limits_namespaces_limitsCmd.Long = organizations_limits_namespaces_limitsCmd.Short
}
