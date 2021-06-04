package scim

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("scim", "SWAGGER_OVERRIDE_/api/v2/scim")
	scimCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scim"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scimCmd)
}

func Cmdscim() *cobra.Command {
	return scimCmd
}
