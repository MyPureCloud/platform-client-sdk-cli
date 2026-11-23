package authorization_customattributes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("authorization_customattributes", "SWAGGER_OVERRIDE_/api/v2/authorization/customattributes")
	authorization_customattributesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_customattributes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_customattributesCmd)
}

func Cmdauthorization_customattributes() *cobra.Command {
	return authorization_customattributesCmd
}
