package audits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("audits", "SWAGGER_OVERRIDE_/api/v2/audits")
	auditsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("audits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(auditsCmd)
}

func Cmdaudits() *cobra.Command {
	return auditsCmd
}
