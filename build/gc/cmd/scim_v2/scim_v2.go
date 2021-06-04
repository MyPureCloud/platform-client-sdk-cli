package scim_v2

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("scim_v2", "SWAGGER_OVERRIDE_/api/v2/scim/v2")
	scim_v2Cmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scim_v2"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scim_v2Cmd)
}

func Cmdscim_v2() *cobra.Command {
	return scim_v2Cmd
}
