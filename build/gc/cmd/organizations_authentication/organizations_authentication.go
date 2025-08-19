package organizations_authentication

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("organizations_authentication", "SWAGGER_OVERRIDE_/api/v2/organizations/authentication")
	organizations_authenticationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_authentication"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_authenticationCmd)
}

func Cmdorganizations_authentication() *cobra.Command {
	return organizations_authenticationCmd
}
