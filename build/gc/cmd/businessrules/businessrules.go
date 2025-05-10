package businessrules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("businessrules", "SWAGGER_OVERRIDE_/api/v2/businessrules")
	businessrulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("businessrules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(businessrulesCmd)
}

func Cmdbusinessrules() *cobra.Command {
	return businessrulesCmd
}
