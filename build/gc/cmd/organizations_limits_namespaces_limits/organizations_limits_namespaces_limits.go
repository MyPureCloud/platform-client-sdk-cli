package organizations_limits_namespaces_limits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("organizations_limits_namespaces_limits", "SWAGGER_OVERRIDE_/api/v2/organizations/limits/namespaces/{namespaceName}/limits")
	organizations_limits_namespaces_limitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_limits_namespaces_limits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_limits_namespaces_limitsCmd)
}

func Cmdorganizations_limits_namespaces_limits() *cobra.Command {
	return organizations_limits_namespaces_limitsCmd
}
