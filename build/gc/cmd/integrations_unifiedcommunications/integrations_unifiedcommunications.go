package integrations_unifiedcommunications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_unifiedcommunications", "SWAGGER_OVERRIDE_/api/v2/integrations/unifiedcommunications")
	integrations_unifiedcommunicationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_unifiedcommunications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_unifiedcommunicationsCmd)
}

func Cmdintegrations_unifiedcommunications() *cobra.Command {
	return integrations_unifiedcommunicationsCmd
}
