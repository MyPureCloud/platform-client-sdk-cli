package orgauthorization_trustor

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("orgauthorization_trustor", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustor")
	orgauthorization_trustorCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustor"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustorCmd)
}

func Cmdorgauthorization_trustor() *cobra.Command {
	return orgauthorization_trustorCmd
}
